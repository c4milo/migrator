// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// +build postgres

package migrator

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var (
	ErrCreatingTable        = errors.New("error-creating-migration-table")
	ErrRegisteringMigration = errors.New("error-registering-migration")
	ErrGettingMigrations    = errors.New("error-getting-migrations")
	ErrRollbackFailed       = errors.New("error-rolling-back-migrations")
	ErrUpdatingMigration    = errors.New("error-updating-migration-metadata")
	ErrRedoFailed           = errors.New("redo-error")
	ErrMigrationNotFound    = errors.New("migration-not-found")
	ErrMigrationIDrequired  = errors.New("migration-id-required")
	ErrDownFailed           = errors.New("migration-down-failed")
)

type postgres struct {
	db           *sql.DB
	paths        []string
	assetFunc    AssetFunc
	assetDirFunc AssetDirFunc
}

func NewPostgres(db *sql.DB, paths []string, assetFunc AssetFunc) (*postgres, error) {
	return &postgres{
		db:        db,
		paths:     paths,
		assetFunc: assetFunc,
	}, nil
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
			-- timestamp of when the migration was created.
			created_at    timestamptz not null default current_timestamp,
			-- timestamp of when the migration was updated.
			updated_at    timestamptz not null,

			primary key (id)
		);
	`)

	if err != nil {
		debug.PrintStack()
		log.Printf("[ERROR] %#v", err)
		return ErrCreatingTable
	}

	return nil
}

// Up re-applies the specific migration ID only if the migration exists and has
// status "down"
func (p *postgres) Up(id string) error {
	ms, err := p.Migrations(id)
	exists := len(ms) > 0
	if !exists {
		return ErrMigrationNotFound
	}

	m := ms[0]

	if m.Status == "up" {
		return nil
	}

	tx, err := p.db.Begin()
	if err != nil {
		debug.PrintStack()
		log.Printf("[ERROR] %#v", err)
		return ErrMigrationFailed
	}

	if _, err := tx.Exec(`DELETE FROM schema_migrations WHERE id = $1`, id); err != nil {
		tx.Rollback()
		return err
	}

	newM, err := DecodeFile(m.Filename, p.assetFunc)
	if err != nil {
		debug.PrintStack()
		log.Printf("[ERROR] %#v", err)
		tx.Rollback()
		return ErrMigrationFailed
	}

	return p.migrate(newM, tx)
}

func (p *postgres) Down(id string) error {
	if id == "" {
		return ErrMigrationIDrequired
	}

	migrations, err := p.db.Query(`
		SELECT id, down FROM schema_migrations
		WHERE status = 'up'
		AND   id = $1`, id)
	if err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrDownFailed
	}
	defer migrations.Close()

	return p.rollback(migrations)
}

func (p *postgres) rollback(migrations *sql.Rows) error {
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

	if err := migrations.Err(); err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrRollbackFailed
	}
	return nil
}

func (p *postgres) Redo(steps ...uint) error {
	n := uint(1)
	if len(steps) > 0 {
		n = steps[0]
	}

	l := len(p.paths)
	numMigrations := uint(l / 2)

	if n > numMigrations {
		n = numMigrations
	}

	if err := p.Rollback(n); err != nil {
		return err
	}

	if err := p.Migrate(); err != nil {
		return err
	}

	return nil
}

func (p *postgres) Rollback(steps ...uint) error {
	n := uint(1)
	if len(steps) > 0 {
		n = steps[0]
	}

	migrations, err := p.db.Query(`
		SELECT id, down FROM schema_migrations
		WHERE status = 'up'
		ORDER BY id DESC LIMIT $1`, n)
	if err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrRollbackFailed
	}
	defer migrations.Close()

	return p.rollback(migrations)
}

func (p *postgres) Migrate() error {
	for i := 0; i < len(p.paths); i++ {
		f := p.paths[i]

		if strings.HasSuffix(f, "down.sql") {
			continue
		}

		m, err := DecodeFile(f, p.assetFunc)
		if err != nil {
			return err
		}

		if err := p.migrate(m, nil); err != nil {
			return err
		}
	}
	return nil
}

func (p *postgres) migrate(m *Migration, currTx *sql.Tx) error {
	tx := currTx
	if tx == nil {
		var err error
		tx, err = p.db.Begin()
		if err != nil {
			log.Printf("[ERROR] %#v", err)
			return ErrMigrationFailed
		}
	}

	ms, err := p.Migrations(m.ID)
	if err != nil {
		log.Printf("[ERROR] %#v", err)
		return ErrMigrationFailed
	}

	exists := len(ms) > 0

	if exists && ms[0].Status == "up" {
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
				id, name, filename, up, down, status, created_at, updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, now(), $7);
		`, m.ID, m.Name, m.Filename, m.Up, m.Down, "up", m.UpdatedAt); err != nil {
			log.Printf("[ERROR] %#v", err)
			tx.Rollback()
			return ErrRegisteringMigration
		}
	} else {
		if _, err := tx.Exec(`
			UPDATE schema_migrations
			SET    status = $1, up = $2, down = $3, updated_at = now()
			WHERE  id = $4`, "up", m.Up, m.Down, m.ID); err != nil {
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

func (p *postgres) Migrations(IDs ...string) ([]*Migration, error) {
	query := `
		SELECT id, name, filename, up, down, status, created_at, updated_at
		FROM schema_migrations
	`

	hasIDs := len(IDs) > 0
	if hasIDs {
		query += ` WHERE id IN (`
		for i := range IDs {
			query += fmt.Sprintf("$%d,", i+1)
		}
		query = strings.TrimSuffix(query, ",")
		query += `)`
	}

	query += ` ORDER by id DESC `

	var err error
	var rows *sql.Rows
	if hasIDs {
		// This is so unfortunate
		new := make([]interface{}, len(IDs))
		for i, v := range IDs {
			new[i] = interface{}(v)
		}
		//log.Printf("%s", query)
		rows, err = p.db.Query(query, new...)
	} else {
		rows, err = p.db.Query(query)
	}

	if err != nil {
		debug.PrintStack()
		log.Printf("[ERROR] %#v", err)
		return nil, ErrGettingMigrations
	}
	defer rows.Close()

	migrations := make([]*Migration, 0)
	for rows.Next() {
		var id, name, filename, up, down, status string
		var createdAt, updatedAt time.Time

		rows.Scan(&id, &name, &filename, &up, &down, &status, &createdAt, &updatedAt)
		m := new(Migration)
		m.ID = id
		m.Name = name
		m.Filename = filename
		m.Up = up
		m.Down = down
		m.Status = status
		m.CreatedAt = createdAt
		m.UpdatedAt = updatedAt

		migrations = append(migrations, m)
	}

	if err = rows.Err(); err != nil {
		debug.PrintStack()
		log.Printf("[ERROR] %#v", err)
		return nil, ErrGettingMigrations
	}
	return migrations, nil
}
