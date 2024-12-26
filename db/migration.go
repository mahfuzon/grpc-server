package db

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func Migrate(conn *sql.DB) {
	log.Println("Migrating database running")
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatalf("error creating postgres driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "bank_db", driver)
	if err != nil {
		log.Fatalf("error creating migration: %v", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("error running migration: %v", err)
	}
}
