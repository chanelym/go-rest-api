package main

import (
	"fmt"
)

func main() {
	fmt.Println("My very first API in Golang <3")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

// Function responsible for instantiating and
// starting up the application.
func Run() error {
	fmt.Println("Starting up the application...")
	return nil
}
