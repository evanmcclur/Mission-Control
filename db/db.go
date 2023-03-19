package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Generate connection string and attempt connection to db
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	return db

	// Read migrations from migrations folder
	// migrations := &migrate.FileMigrationSource{
	// 	Dir: "migrations/postgres",
	// }

	// n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	// if err != nil {
	// 	// Handle errors!
	// 	log.Fatalf("Error applying migrations to database: %s", err)
	// }
	// fmt.Printf("Applied %d migrations!\n", n)
}
