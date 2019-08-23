package api

import (
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
	Router *mux.Router
	DB     Database
}

func (a *App) initializeRoutes() {
	fmt.Println("Initialize the app routes")
	a.Router.HandleFunc("/", RootHandler)

	a.Router.HandleFunc("/books", a.ListBooksHandler).Methods("GET")
	a.Router.HandleFunc("/books/{id:[0-9]+}", a.GetBookHandler).Methods("GET")

	// a.Router.HandleFunc("/books", a.createBook).Methods("POST")
	// a.Router.HandleFunc("/books/{id:[0-9]+}", a.updateBook).Methods("PUT")
	// a.Router.HandleFunc("/books/{id:[0-9]+}", a.deleteBook).Methods("DELETE")
}

// Initialize the app
func (a *App) Initialize() error {
	a.Router = mux.NewRouter()
	if err := a.DB.seedData(); err != nil {
		return err
	}
	a.initializeRoutes()
	return nil
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
