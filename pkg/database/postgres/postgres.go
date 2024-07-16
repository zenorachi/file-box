package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	migrationDir = "./migrations"
	driverName   = "postgres"
)

type DBConfig struct {
	DSN            string
	MigrationTable string
	MaxIdleConns   int
	MaxOpenConns   int
	AutoMigrate    bool
}

func NewDB(cfg *DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driverName, cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxIdleConns)

	if err = goose.SetDialect(driverName); err != nil {
		return nil, err
	}

	goose.SetTableName(cfg.MigrationTable)
	if !cfg.AutoMigrate {
		return db, nil
	}

	if err = goose.Up(db.DB, migrationDir); err != nil {
		return nil, err
	}

	return db, nil
}
