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
	sdcj := jobs.SpaceDotComJob{LastRun: nil}
	cm := cron.CronManager{Crons: map[cron.JobKey]cron.CronJob{}}
	cm.RegisterJob(sdcj)

	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(sdcj.String())
	})
	_ = http.ListenAndServe(":8080", nil)
}
