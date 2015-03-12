// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package migrator

import (
	"database/sql"
	"errors"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	// ErrBadFilenameFormat is returned if a migration file does not conform
	// with the expected name patter. Ex: 001_my-migration_up.sql
	ErrBadFilenameFormat = errors.New("bad-filename-format")
	// ErrInvalidDirection is returned if the direction of a migration is other
	// than 'up' or 'down'
	ErrInvalidDirection = errors.New("invalid-migration-direction")
	// ErrDBNotSupported is returned when trying to create a migrator instance
	// for an unsupported database.
	ErrDBNotSupported = errors.New("database-not-supported")
	// ErrInvalidDB is returned when a nil *sql.DB pointer is passed to NewMigrator
	ErrInvalidDB = errors.New("invalid-database-handle")

	ErrMigrationFailed = errors.New("migration-failed")
)

type DBType string

// Supported databases.
const (
	Postgres DBType = "postgres"
)

type AssetFunc func(path string) ([]byte, error)
type AssetDirFunc func(path string) ([]string, error)

type Migrator interface {
	// Init initializes migrations table in the given database.
	Init() error
	// Migrate applies all migrations that hasn't been applied.
	Migrate() error
	// Redo undos specific migrations and applies them again. By default
	// if not parameter is specified, it will redo the latest migration.
	Redo(n ...uint) error
	// Rollback reverts the last migration if not parameter is specified.
	Rollback(n ...uint) error
	// Migrations returns the list of migrations currently applied to the database.
	Migrations(ids ...string) ([]*Migration, error)
	// Up applies a specific migration version.
	Up(version string) error
	// Down rolls back or takes down a specific migration version.
	Down(version string) error
}

// Migration represents an actual migration file.
type Migration struct {
	ID        string
	Name      string
	Filename  string `db:"filename"`
	Up        string
	Down      string
	Status    string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const baseDir string = "migrations/postgres"

func NewMigrator(db *sql.DB, dbType DBType, assetFunc AssetFunc, assetDirFunc AssetDirFunc) (Migrator, error) {
	if db == nil {
		return nil, ErrInvalidDB
	}

	paths, err := assetDirFunc(baseDir)
	if err != nil {
		log.Printf("[ERROR] trying to get list of migration files from embedded asset directory %s", baseDir)
		log.Printf("[ERROR] %#v", err)
		return nil, ErrMigrationFailed
	}

	sort.Strings(paths)

	var migrator Migrator
	switch dbType {
	case Postgres:
		var err error
		migrator, err = NewPostgres(db, paths, assetFunc)
		if err != nil {
			return nil, err
		}
	}

	if migrator == nil {
		return nil, ErrDBNotSupported
	}

	if err := migrator.Init(); err != nil {
		return nil, err
	}

	return migrator, nil
}

// DecodeFile takes a sql file and returns a Migration instance
func DecodeFile(f string, assetFunc AssetFunc) (*Migration, error) {
	// File names should be formatted like so: id_migration-name_up.sql or
	// id_migration-name_down.sql. Ex: 0002_create-extension-citext_down.sql
	parts := strings.Split(f, "_")
	if len(parts) != 3 {
		log.Printf("[ERROR] Bad file format: %s", f)
		return nil, ErrBadFilenameFormat
	}

	m := new(Migration)
	m.ID = parts[0]
	m.Name = parts[1]
	m.Filename = f

	upFile := filepath.Join(baseDir, f)
	upSQL, err := assetFunc(upFile)
	if err != nil {
		log.Printf("[ERROR] Extracting asset content from %s", upFile)
		log.Printf("[ERROR] %#v", err)
		return nil, ErrMigrationFailed
	}

	downFile := filepath.Join(baseDir, strings.Replace(f, "up.sql", "down.sql", 1))
	downSQL, err := assetFunc(downFile)
	if err != nil {
		log.Printf("[ERROR] Extracting asset content from %s", downFile)
		log.Printf("[ERROR] %#v", err)
		return nil, ErrMigrationFailed
	}

	m.Up = string(upSQL[:])
	m.Down = string(downSQL[:])
	return m, nil
}
