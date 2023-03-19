package main

import (
	"fmt"
	"net/http"

	// "github.com/go-co-op/gocron"
	"github.com/evanmcclur/mission-control/db"
	"github.com/evanmcclur/mission-control/handlers"
)

func main() {
	db.InitDB()
	router := handlers.GetRouter()

	fmt.Println("Listening on port 8080...")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Hello, World")
	// })
	_ = http.ListenAndServe(":8080", router)
}
