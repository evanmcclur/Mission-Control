package cron

type Status uint

// Define status constants
const (
	Success Status = 0
	Failed  Status = 1
)

type JobKey uint

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
