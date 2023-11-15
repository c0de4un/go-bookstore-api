package main

import (
	"fmt"
	"net/http"
)

type Router struct {
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "{}")
}

func main() {
	fmt.Print("Starting server on :3000 ...")
	var router Router
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
