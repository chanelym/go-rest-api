package main

import (
	"context"
	"fmt"

	book "github.com/chanelym/go-rest-api/internal/books"
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
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Fail to migrate database")
		return err
	}

	bookService := book.NewService(db)
	fmt.Println(bookService.GetBook(
		context.Background(),
		"9a31bf83-28dc-4b8d-bf70-7d347a24ff2e",
	))

	return nil
}
