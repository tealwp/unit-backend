package cmd

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"

	"github.com/tealwp/unit-backend/pkg/cfg"
	"github.com/tealwp/unit-backend/pkg/db/migration"
)

func Migrate(cfg *cfg.Config) error {
	flag.Parse()

	if len(flag.Args()) < 2 {
		return errors.New("no subcommand for migration")
	}

	dbConn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return fmt.Errorf("creating db instance: %v", err)
	}

	subcommand := flag.Args()[1]
	switch subcommand {
	case "up":
		err := migration.MigrateUp(db)
		if err != nil {
			return err
		}
	case "down":
		err := migration.MigrateDown(db)
		if err != nil {
			return err
		}
	case "ping":
		err := migration.Ping(db)
		if err != nil {
			return err
		}
		fmt.Println("Database connection verified.")
	default:
		return fmt.Errorf("migration subcommand: %s is not a proper input", subcommand)
	}
	return nil
}
