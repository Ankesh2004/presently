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

func (w *Book) getBook(db *sql.DB) error {
	return db.QueryRow("SELECT id, title, author, publish_year from books WHERE id=$1",
		w.ID).Scan(&w.ID, &w.Title, &w.Author, &w.PublishYear)
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
