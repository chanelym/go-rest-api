package book

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingBook   = errors.New("failed to fetch book by id")
	ErrNotImplemented = errors.New("not implemented")
)

// Comment - representation of the book structure.
type Book struct {
	ID     string
	Path   string
	Title  string
	Author string
}

/*Store - Here I'll put all the methods that my service needs
to operate. By doing that, it gets more easily to unit test everything*/
type Store interface {
	GetBook(context.Context, string) (Book, error)
}

/* Service - struct on which I'll build all of the logic.
Here I'll receive all of the methods that will have a
pointer to this struct.*/
type Service struct {
	Store Store
}

/* NewService - returns a pointer to a new service.
Making this "constructor function" allow me to add
more complex logic in the future.*/
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
		return Book{}, ErrFetchingBook
	}

	return book, nil
}

func (s *Service) UpdateBook(ctx context.Context, book Book) error {
	return ErrNotImplemented
}

func (s *Service) DeleteBook(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateBook(ctx context.Context, book Book) (Book, error) {
	return Book{}, ErrNotImplemented
}
