// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// +build postgres

package migrator

import (
	"database/sql"
	"errors"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var (
	ErrCreatingTable        = errors.New("error-creating-migration-table")
	ErrMigrationFailed      = errors.New("migration-failed")
	ErrRegisteringMigration = errors.New("error-registering-migration")
	ErrGettingMigrations    = errors.New("error-getting-migrations")
	ErrRollbackFailed       = errors.New("error-rolling-back-migrations")
	ErrUpdatingMigration    = errors.New("error-updating-migration-metadata")
)

type postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *postgres {
	return &postgres{db: db}
}

func (p *postgres) Init() error {
	_, err := p.db.Exec(`
		-- creates an enum type for OAuth2 token types
		do $$
		begin
			if not exists (select 1 from pg_type where typname = 'migration_status_type') then
				create type migration_status_type as enum ('up', 'down');
			end if;
		end$$;
		-- creates citext extension
		create extension if not exists citext;

		-- creates table schema_migrations
		create table if not exists schema_migrations (
			-- migration identifier as found in migration file name.
			id            citext not null,
			-- migration name as found in migration file name.
			name          text   not null,
			-- migration file name.
			filename      text   not null,
			-- migration sql content as found in Up migration file.
			up            text   not null,
			-- migration sql content as found in Down migration file.
			down          text   not null,
			-- status of this migration
			status        migration_status_type,
			-- timestamp of when the grant was created.
			created_at    timestamptz not null default current_timestamp,

			primary key (id)
		);
	`)

	if err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrCreatingTable
	}

	return nil
}

// isApplied checks whether the migration was already applied or not
func (p *postgres) isApplied(id string) (applied, exist bool, err error) {
	var status string
	row := p.db.QueryRow(`SELECT status FROM schema_migrations WHERE id = $1`, id)
	err = row.Scan(&status)
	if err != nil {
		if sql.ErrNoRows == err {
			return false, false, nil
		}

		return false, false, err
	}

	if status == "up" {
		return true, true, nil
	}

	return false, true, nil
}

func (p *postgres) MigrateFromAsset(assetFunc AssetFunc, assetDirFunc AssetDirFunc) error {
	baseDir := "migrations/postgres"
	files, err := assetDirFunc(baseDir)
	if err != nil {
		log.Printf("[ERROR] trying to get list of migration files from embedded asset directory %s", baseDir)
		log.Printf("[ERROR] %#v", err)
		return ErrMigrationFailed
	}

	sort.Strings(files)

	for i := 0; i < len(files); i++ {
		f := files[i]
		// File names should be formatted like so: id_migration-name_up.sql or
		// id_migration-name_down.sql. Ex: 0002_create-extension-citext_down.sql
		parts := strings.Split(f, "_")
		if len(parts) != 3 {
			log.Printf("[ERROR] Bad file format: %s", f)
			return ErrBadFilenameFormat
		}

		m := new(Migration)
		m.ID = parts[0]
		m.Name = parts[1]
		m.Filename = f

		if parts[2] != "up.sql" {
			continue
		}

		upFile := filepath.Join(baseDir, f)
		upSQL, err := assetFunc(upFile)
		if err != nil {
			log.Printf("[ERROR] Extracting asset content from %s", upFile)
			log.Printf("[ERROR] %#v", err)
			return ErrMigrationFailed
		}

		downFile := filepath.Join(baseDir, strings.Replace(f, "up.sql", "down.sql", 1))
		downSQL, err := assetFunc(downFile)
		if err != nil {
			log.Printf("[ERROR] Extracting asset content from %s", downFile)
			log.Printf("[ERROR] %#v", err)
			return ErrMigrationFailed
		}

		m.Up = string(upSQL[:])
		m.Down = string(downSQL[:])

		if err := p.migrate(m); err != nil {
			return err
		}
	}
	return nil
}

func (p *postgres) migrate(m *Migration) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrMigrationFailed
	}

	applied, exists, err := p.isApplied(m.ID)
	if err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrMigrationFailed
	}

	if applied {
		return nil
	}

	if _, err := tx.Exec(m.Up); err != nil {
		log.Printf("[ERROR] %#v", err)
		log.Printf("[ERROR] %s", m.Up)
		tx.Rollback()
		return ErrMigrationFailed
	}

	if !exists {
		if _, err := tx.Exec(`
			INSERT INTO schema_migrations (
				id, name, filename, up, down, status, created_at
			) VALUES ($1, $2, $3, $4, $5, $6, now());
		`, m.ID, m.Name, m.Filename, m.Up, m.Down, "up"); err != nil {
			log.Printf("[ERROR] %#v", err)
			tx.Rollback()
			return ErrRegisteringMigration
		}
	} else {
		if _, err := tx.Exec(`
			UPDATE schema_migrations
			SET    status = $1
			WHERE  id = $2`, "up", m.ID); err != nil {
			log.Printf("[ERROR] %#v", err)
			tx.Rollback()
			return ErrUpdatingMigration
		}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("[ERROR] %#v", err)
		tx.Rollback()
		return ErrMigrationFailed
	}

	return nil
}

func (p *postgres) Rollback() error {
	return p.RollbackN(1)
}

func (p *postgres) RollbackN(n uint) error {
	migrations, err := p.db.Query(`
		SELECT id, down FROM schema_migrations
		WHERE status = 'up'
		ORDER BY id DESC LIMIT $1`, n)
	if err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrRollbackFailed
	}
	defer migrations.Close()

	for migrations.Next() {
		var id, downSQL string
		if err := migrations.Scan(&id, &downSQL); err != nil {
			log.Printf("[ERROR] %#v", err)
			return ErrRollbackFailed
		}

		tx, err := p.db.Begin()
		if err != nil {
			log.Printf("[ERROR] %#v", err)
			return ErrRollbackFailed
		}

		if _, err = tx.Exec(downSQL); err != nil {
			log.Printf("[ERROR] %#v", err)
			log.Printf("[ERROR] Down query: %s", downSQL)
			tx.Rollback()
			return ErrRollbackFailed
		}

		if _, err = tx.Exec(`
			UPDATE schema_migrations
			SET    status = $1
			WHERE  id = $2
		`, "down", id); err != nil {
			log.Printf("[ERROR] id=%s status=%s err=%#v", id, "down", err)
			tx.Rollback()
			return ErrRollbackFailed
		}

		if err := tx.Commit(); err != nil {
			log.Printf("[ERROR] %#v", err)
			tx.Rollback()
			return ErrRollbackFailed
		}
	}

	if err = migrations.Err(); err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrRollbackFailed
	}
	return nil
}

func (p *postgres) Migrations() ([]*Migration, error) {
	rows, err := p.db.Query(`
		SELECT id, name, filename, up, down, status, created_at
		FROM schema_migrations
	`)
	if err != nil {
		log.Printf("[ERROR] %#v", err)
		return nil, ErrGettingMigrations
	}
	defer rows.Close()
	migrations := make([]*Migration, 0)
	for rows.Next() {
		var id, name, filename, up, down, status string
		var createdAt time.Time

		rows.Scan(&id, &name, &filename, &up, &down, &status, &createdAt)
		m := new(Migration)
		m.ID = id
		m.Name = name
		m.Filename = filename
		m.Up = up
		m.Down = down
		m.Status = status
		m.CreatedAt = createdAt

		migrations = append(migrations, m)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[ERROR] %#v", err)
		return nil, ErrGettingMigrations
	}
	return migrations, nil
}
