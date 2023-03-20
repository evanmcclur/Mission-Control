package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

// Singleton object that encapsulates sql.DB object. This ensures that only ONE that one long-lived sql.DB object is created.
type DBSingleton struct {
	Context *sql.DB
}

var instance *DBSingleton
var once *sync.Once = new(sync.Once)

func GetInstance() *DBSingleton {
	once.Do(func() {
		db, err := initDB()
		if err != nil {
			log.Fatal(err)
		}
		instance = &DBSingleton{
			Context: db,
		}
		fmt.Println("Created new instance!")
	})

	return instance
}

// Use this instead of calling .Close() on sql.DB object.
// Calling this will close the database and destroys the DBSingleton instance. This allows a new Singleton instance to be created after one is destroyed.
func (db *DBSingleton) Destroy() error {
	err := instance.Context.Close()
	if err != nil {
		return err
	}
	once = new(sync.Once)
	instance = nil
	return nil
}

// Connects the database and ensures that it is available for use during program execution. This should only be called once.
func initDB() (*sql.DB, error) {
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
		return nil, err
	}

	// Ensure the db is available for use
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
