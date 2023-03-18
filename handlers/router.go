package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/articles", GetAllArticles).Methods(http.MethodGet)
	// TODO: Restrict to admins only
	r.HandleFunc("/jobs", GetAllJobs).Methods(http.MethodGet)

	return r
}
