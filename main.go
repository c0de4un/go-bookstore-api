package main

import (
	"fmt"
	"net/http"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "{}")
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Print("Starting server on :3000 ...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
