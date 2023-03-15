package cron

import (
	"time"
)

// This cron is used to scrape articles off of space.com
type SpaceDotComCron struct {
	// When this cron job was last run
	LastRun *time.Time
}

func (s SpaceDotComCron) Run() (Status, error) {
	status := Success

	return status, nil
}

func (s SpaceDotComCron) Key() JobKey {
	return SDC
}

func (s SpaceDotComCron) String() string {
	return "space.com"
}
