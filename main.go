package main

import (
	"fmt"
	"net/http"

	// "github.com/go-co-op/gocron"
	"github.com/evanmcclur/mission-control/cron"
	"github.com/evanmcclur/mission-control/cron/jobs"
)

func main() {
	// example use of cron manager
	sdcj := &jobs.SpaceDotComJob{}
	cm := cron.CronManager{Crons: map[cron.JobKey]cron.CronJob{}}
	cm.RegisterJob(sdcj)

	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		status, _ := sdcj.Run()
		fmt.Println(status)
	})
	_ = http.ListenAndServe(":8080", nil)
}
