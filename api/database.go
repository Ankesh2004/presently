package api

import (
	"database/sql"
	"fmt"
)

type Database interface {
	getBook(id int) (*Book, error)
	listBooks(start, count int) ([]Book, error)
	seedData() error
}

type SqliteDatabase struct {
	DB *sql.DB
}

func CreateSqliteDatabase() (*SqliteDatabase, error) {
	sqldb, err := sql.Open("sqlite3", "../books.db")
	if err != nil {
		return nil, err
	}
	return &SqliteDatabase{DB: sqldb}, nil
}

func (sb SqliteDatabase) getBook(id int) (*Book, error) {
	book := Book{ID: id}
	err := sb.DB.QueryRow("SELECT title, author, publish_year from books WHERE id=$1", id).Scan(&book.Title, &book.Author, &book.PublishYear)

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

func (sb SqliteDatabase) listBooks(start, count int) ([]Book, error) {
	rows, err := sb.DB.Query(
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

func (sq SqliteDatabase) seedData() error {
	fmt.Println("Ensure the books table exists")
	statement, _ := sq.DB.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, publish_year NUMERIC)")
	if _, err := statement.Exec(); err != nil {
		return err
	}

	var seedBooks = SeedData()
	fmt.Printf("Seed %d book rows in database\n", len(seedBooks))

	// Insert the test data
	for _, book := range seedBooks {
		fmt.Printf("Insert %d into books table\n", book.ID)
		statement, _ := sq.DB.Prepare("INSERT INTO books (id, title, author, publish_year) VALUES (?, ?, ?, ?)")
		if _, err := statement.Exec(book.ID, book.Title, book.Author, book.PublishYear); err != nil {
			return err
		}
	}
	fmt.Println("Done seeding the database")
	return nil
}
