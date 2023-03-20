package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/go-co-op/gocron"
	"github.com/evanmcclur/mission-control/db"
	"github.com/evanmcclur/mission-control/db/migrations"
	"github.com/evanmcclur/mission-control/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := db.GetInstance()
	// Apply latest migrations after establishing connection with database
	err = migrations.ApplyMigrations(ctx)
	if err != nil {
		log.Fatalf("Error applying migrations to database: %s", err)
	}

	router := handlers.GetRouter()

	fmt.Println("Listening on port 8080...")
	_ = http.ListenAndServe(":8080", router)

	// Destroy instance and close database...
	ctx.Destroy()
}
