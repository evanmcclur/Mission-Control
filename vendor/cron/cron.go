package cron

type Status uint

// Define status constants
const (
	Success Status = 0
	Failed  Status = 1
)

// A unique key that identifies a particular cron job
type JobKey uint

// Keys that represent their respective cron job. When adding a new cron job, update this.
const (
	SDC JobKey = 1
)

// Defines the implementation for all cron jobs
type CronJob interface {
	Run() (Status, error)
	Key() JobKey
	String() string
}

// Converts status from constant to string form
func (status Status) String() string {
	statuses := [...]string{
		"SUCCESS",
		"FAILED",
	}

	if status < Success || status > Failed {
		return "UNKNOWN"
	}

	return statuses[status]
}
