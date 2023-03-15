package main

import (
	"fmt"
	"net/http"
	// "github.com/go-co-op/gocron"
)

// func hello(name string) {
// 	message := fmt.Sprintf("Hi, %v", name)
// 	fmt.Println(message)
// }

// func runCronJobs() {
// 	s := gocron.NewScheduler(time.UTC)

// 	s.Every(1).Seconds().Do(func() {
// 		hello("Selma stinks")
// 	})

// 	s.StartBlocking()
// 	s.Stop()
// }

func main() {
	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Hello World")
	})
	_ = http.ListenAndServe(":8080", nil)
}
