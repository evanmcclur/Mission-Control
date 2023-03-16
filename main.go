package main

import (
	"fmt"
	"net/http"

	// "github.com/go-co-op/gocron"
	"cron"
)

func main() {
	status := cron.Success
	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(status.String())
	})
	_ = http.ListenAndServe(":8080", nil)
}
