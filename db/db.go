package db

import (
	"database/sql"
	"fmt"

	"github.com/rubenv/sql-migrate"
)

func InitDB() {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		// Handle errors!
	}

	// Read migrations from migrations folder
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/postgres",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
