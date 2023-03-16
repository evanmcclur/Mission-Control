package jobs

import (
	"time"

	"cron"
)

// This cron is used to scrape articles off of space.com
type SpaceDotComJob struct {
	// When this cron job was last run
	LastRun *time.Time
}

func (s SpaceDotComJob) Run() (cron.Status, error) {
	status := cron.Success

	return status, nil
}

func (s SpaceDotComJob) Key() cron.JobKey {
	return cron.SDC
}

func (s SpaceDotComJob) String() string {
	return "space.com"
}
