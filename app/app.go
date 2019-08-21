package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	// Using sqlite provider with database/sql
	_ "github.com/mattn/go-sqlite3"
)

// The App struct represents the core application object
type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (a *App) initializeRoutes() {
	fmt.Println("Initialize the app routes")
	a.Router.HandleFunc("/", RootHandler)

	a.Router.HandleFunc("/hello", HelloHandler).Methods("GET")

	a.Router.HandleFunc("/books", a.ListBooksHandler).Methods("GET")

	// a.Router.HandleFunc("/widget", a.createProduct).Methods("POST")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.getProduct).Methods("GET")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *App) initializeDb() {
	fmt.Println("Ensure the books table exists")
	statement, _ := a.Database.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, publish_year NUMERIC)")
	statement.Exec()

	// var seedBooks = []*Book{
	// 	&Book{
	// 		ID:          1,
	// 		Title:       "Ulysses",
	// 		Author:      "James Joyce",
	// 		PublishYear: 1922,
	// 	},
	// 	&Book{
	// 		ID:          2,
	// 		Title:       "The Great Gatsby",
	// 		Author:      "F Scott Fitzgerald",
	// 		PublishYear: 1925,
	// 	},
	// 	&Book{
	// 		ID:          3,
	// 		Title:       "Moby Dick",
	// 		Author:      "Herman Melville",
	// 		PublishYear: 1851,
	// 	},
	// 	&Book{
	// 		ID:          4,
	// 		Title:       "War and Peace",
	// 		Author:      "Leo Tolstoy",
	// 		PublishYear: 1869,
	// 	},
	// }

	var seedBooks = SeedData()
	fmt.Printf("Seed %d book rows in database\n", len(seedBooks))

	// Insert the test data
	for _, book := range seedBooks {
		fmt.Printf("Insert %d into books table\n", book.ID)
		statement, _ := a.Database.Prepare("INSERT INTO books (id, title, author, publish_year) VALUES (?, ?, ?, ?)")
		statement.Exec(book.ID, book.Title, book.Author, book.PublishYear)
	}
	fmt.Println("Done seeding the database")
}

// Initialize the app
// Parameters for one-time app initialization steps passed in here, i.e. db connection args.
// func (a *App) Initialize(user, password, dbname string) {
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	database, _ := sql.Open("sqlite3", "./books.db")
	a.Database = database

	a.initializeDb()
	a.initializeRoutes()
}

// Run the app
func (a *App) Run(port int) error {
	fmt.Printf("Run the app on port %d\n", port)
	loggedRouter := handlers.LoggingHandler(os.Stdout, a.Router)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)

	// srv := &http.Server{
	// 	Addr: fmt.Sprintf("0.0.0.0:%d", port),
	// 	// Good practice to set timeouts to avoid Slowloris attacks.
	// 	WriteTimeout: time.Second * 15,
	// 	ReadTimeout:  time.Second * 15,
	// 	IdleTimeout:  time.Second * 60,
	// 	Handler:      loggedRouter,
	// }

	// err := srv.ListenAndServe()
	// if err != nil {
	// 	return nil, err
	// }

	// return srv, nil
}
