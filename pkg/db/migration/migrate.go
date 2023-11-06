package migration

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func newMigration(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("creating driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./pkg/db/migration/migrations",
		"postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("creating migration: %v", err)
	}

	return m, nil
}

func MigrateUp(db *sql.DB) error {
	m, err := newMigration(db)
	if err != nil {
		return fmt.Errorf("creating new migration: %v", err)
	}
	m.Up()
	fmt.Println("migrated up successfully")
	return nil
}

func MigrateDown(db *sql.DB) error {
	m, err := newMigration(db)
	if err != nil {
		return fmt.Errorf("creating new migration: %v", err)
	}
	m.Down()
	fmt.Println("migrated down successfully")
	return nil
}

func Ping(db *sql.DB) error {
	err := db.Ping()
	return err
}
