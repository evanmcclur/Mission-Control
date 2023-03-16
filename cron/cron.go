package cron

// Represents the result of running a cron job
type Status uint

// Status constants that represent the status of a job after it runs
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

func (jk JobKey) String() string {
	// NOTE: update this whenever adding a new cron job
	keys := [...]string{
		"SDC",
	}

	return keys[jk]
}
