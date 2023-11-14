package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World !</h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Print("Starting server on :3000 ...")
	http.ListenAndServe(":3000", nil)
}
