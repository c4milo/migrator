package migrator

import (
	"database/sql"
	"errors"
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
	// Migrate runs a migration from an embedded asset. Go-bindata is the only
	// library supported.
	MigrateFromAsset(assetFunc AssetFunc, assetDirFunc AssetDirFunc) error
	// Rollback reverts the last migration.
	Rollback() error
	// RollbackN reverts the last N migrations.
	RollbackN(n uint) error
	// Migrations returns the list of migrations currently applied to the database.
	Migrations() ([]*Migration, error)
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
}

func NewMigrator(db *sql.DB, dbType DBType) (Migrator, error) {
	var migrator Migrator

	switch dbType {
	case Postgres:
		migrator = NewPostgres(db)
	}

	if migrator == nil {
		return nil, ErrDBNotSupported
	}

	if err := migrator.Init(); err != nil {
		return nil, err
	}

	return migrator, nil
}
