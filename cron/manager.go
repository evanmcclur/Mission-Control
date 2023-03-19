package cron

// import (
// 	"github.com/go-co-op/gocron"
// )

type CronManager struct {
	// Registered cron jobs
	Crons map[JobKey]CronJob
}

// Constructs a new CronManager
func NewCronManager() *CronManager {
	return &CronManager{
		Crons: make(map[JobKey]CronJob),
	}
}

// Get's all registered jobs from the manager
func (cm *CronManager) GetAllJobs() []CronJob {
	jobs := make([]CronJob, 0, len(cm.Crons))
	for jk := range cm.Crons {
		jobs = append(jobs, cm.Crons[jk])
	}

	return jobs
}

// Registers a new job with the cron manager
func (cm *CronManager) RegisterJob(job CronJob) {
	key := job.Key()
	cm.Crons[key] = job
}

// Removes a job from the cron manager
func (cm *CronManager) RemoveJob(job CronJob) {
	key := job.Key()
	delete(cm.Crons, key)
}
