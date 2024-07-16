package postgres

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const driverName = "postgres"

// DBConfig - configuration for the database.
// DSN - full database url (including user, password, etc.).
// MigrationDir - path to the migration directory (default: "./migrations").
// MigrationTable - name of the migration table.
// MaxIdleConns - maximum number of idle (default: 100).
// MaxOpenConns - maximum number of open connections (default: 10).
// AutoMigrate - if true, will run auto migration (default: false).
type DBConfig struct {
	DSN            string `required:"true"`
	MigrationTable string `required:"true"  split_words:"true"`
	MigrationDir   string `required:"false" split_words:"true" default:"./migrations"`
	MaxIdleConns   int    `required:"false" split_words:"true" default:"100"`
	MaxOpenConns   int    `required:"false" split_words:"true" default:"10"`
	AutoMigrate    bool   `required:"false" split_words:"true" default:"false"`
}

// NewDB creates a new DB connection using the given config DBConfig.
// Also runs auto migration if AutoMigrate in config value is true. Default path for migration directory is "./migrations".
func NewDB(cfg *DBConfig) (*sqlx.DB, error) {
	if cfg == nil {
		return nil, errors.New("postgres config is required")
	}

	db, err := sqlx.Connect(driverName, cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxIdleConns)

	if !cfg.AutoMigrate {
		return db, nil
	}

	if err = goose.SetDialect(driverName); err != nil {
		return nil, err
	}

	goose.SetTableName(cfg.MigrationTable)
	if err = goose.Up(db.DB, cfg.MigrationDir); err != nil {
		return nil, err
	}

	return db, nil
}