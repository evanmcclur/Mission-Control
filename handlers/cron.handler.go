package handlers

import (
	"net/http"
)

// This handler should be restricted to ADMIN users only

func GetAllJobs(w http.ResponseWriter, r *http.Request) {
	// TODO: Gets all currently scheduled cron jobs
}
