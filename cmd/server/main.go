package main

import (
	"context"
	"fmt"

	"github.com/chanelym/go-rest-api/internal/database"
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

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Fail to connect to database: ")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("database sucessfully connected and pinged")

	return nil
}
