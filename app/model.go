// model.go

package app

import (
	"database/sql"
)

// The Book model
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishYear int    `json:"publish_year"`
}

func getBook(db *sql.DB, id int) (*Book, error) {
	book := Book{ID: id}
	err := db.QueryRow("SELECT title, author, publish_year from books WHERE id=$1", id).Scan(&book.Title, &book.Author, &book.PublishYear)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	return &book, nil
}

func listBooks(db *sql.DB, start, count int) ([]Book, error) {
	rows, err := db.Query(
		"SELECT id, title, author, publish_year FROM books LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishYear); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
