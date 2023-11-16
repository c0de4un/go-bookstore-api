package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	host string
}

func NewAPIServer(host string) *APIServer {
	return &APIServer{
		host: host,
	}
}

func (server *APIServer) Run() error {
	router := chi.NewRouter()
	router.NotFound(server.handleNotFound)

	return http.ListenAndServe(":3000", router)
}

func (server *APIServer) handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "{}")
}
