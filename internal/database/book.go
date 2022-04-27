package database

import (
	"context"
	"database/sql"
	"fmt"

	book "github.com/chanelym/go-rest-api/internal/books"
)

type BookRow struct {
	ID     string
	Path   sql.NullString
	Title  sql.NullString
	Author sql.NullString
}

func convertBookRowToBook(b BookRow) book.Book {
	return book.Book{
		ID:     b.ID,
		Path:   b.Path.String,
		Title:  b.Title.String,
		Author: b.Author.String,
	}
}

func (d *Database) GetBook(
	ctx context.Context,
	uuid string,
) (book.Book, error) {
	var bookRow BookRow
	row := d.Client.QueryRowxContext(
		ctx,
		`SELECT id, path, title, author
		FROM books
		WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&bookRow.ID, &bookRow.Path, &bookRow.Title, &bookRow.Author)
	if err != nil {
		return book.Book{}, fmt.Errorf("error fetching the book with uuid: %w", err)
	}

	return convertBookRowToBook(bookRow), nil
}
