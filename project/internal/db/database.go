package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct{}

func (d Database) Connect(cfg Config) error {
	db, err := sql.Open("mysql", cfg.DNS())
	if err != nil {
		return err
	}

	return d.ping(db)
}

func (d Database) ping(conn *sql.DB) error {
	return conn.Ping()
}
