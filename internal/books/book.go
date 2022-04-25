package book

import (
	"context"
	"fmt"
)

type Store interface {
	GetBook(context.Context, string) (Book, error)
}

// Comment - representation of the book structure
type Book struct {
	ID     string
	Slug   string
	Title  string
	Author string
}

// Service - struct on which I'll build
// all of the logic
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetBook(ctx context.Context, id string) (Book, error) {
	fmt.Println("Retrieving Book information...")
	book, err := s.Store.GetBook(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Book{}, err
	}

	return book, nil
}
