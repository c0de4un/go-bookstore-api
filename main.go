package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "{}")
}

func main() {
	router := chi.NewRouter()
	router.NotFound(handleNotFound)

	fmt.Print("Starting server on :3000 ...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
