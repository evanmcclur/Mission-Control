package migrations

import (
	"fmt"

	"github.com/evanmcclur/mission-control/db"
	migrate "github.com/rubenv/sql-migrate"
)

func ApplyMigrations(db *db.DBSingleton) error {
	// Read migrations from migrations folder
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations/postgres",
	}

	n, err := migrate.Exec(db.Context, "postgres", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
		return err
		// log.Fatalf("Error applying migrations to database: %s", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
