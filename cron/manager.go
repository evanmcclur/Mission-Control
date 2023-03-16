package cron

type CronManager struct {
	// Registered cron jobs
	Crons map[JobKey]CronJob
}

func (cm *CronManager) RegisterJob(job CronJob) {
	key := job.Key()
	cm.Crons[key] = job
}

func (cm *CronManager) RemoveJob(job CronJob) {
	key := job.Key()
	delete(cm.Crons, key)
}
