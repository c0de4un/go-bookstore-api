package main

import (
	"fmt"
)

func main() {
	fmt.Print("Starting server on :3000 ...")

	err := NewAPIServer(":3000").Run()
	if err != nil {
		panic(err)
	}
}
