package app

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	a.Router.HandleFunc("/", RootHandler)

	a.Router.HandleFunc("/hello", HelloHandler).Methods("GET")

	a.Router.HandleFunc("/widgets", a.ListWidgetsHandler).Methods("GET")
	// a.Router.HandleFunc("/widget", a.createProduct).Methods("POST")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.getProduct).Methods("GET")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *App) initializeDb() {
	statement, _ := a.Database.Prepare("CREATE TABLE IF NOT EXISTS widgets (id INTEGER PRIMARY KEY, name TEXT, price NUMERIC, description TEXT)")
	statement.Exec()

	var testWidgets = []*Widget{
		&Widget{
			ID:          1,
			Name:        "Rubber Mallet",
			Price:       19.95,
			Description: "Your basic mallet for hitting with",
		},
		&Widget{
			ID:          2,
			Name:        "Allen Wrench",
			Price:       15,
			Description: "Needed for assembling IKEA furniture",
		},
	}

	fmt.Printf("Insert %d test widgets", len(testWidgets))

	// Insert the test data
	for _, w := range testWidgets {
		fmt.Printf("Insert test row %d into widgets table\n", w.ID)
		statement, _ := a.Database.Prepare("INSERT INTO widgets (id, name, price, description) VALUES (?, ?, ?, ?)")
		statement.Exec(w.ID, w.Name, w.Price, w.Description)
	}
}

// Initialize the app
// Parameters for one-time app initialization steps passed in here, i.e. db connection args.
// func (a *App) Initialize(user, password, dbname string) {
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	database, _ := sql.Open("sqlite3", "./widgets.db")
	a.Database = database

	a.initializeDb()
	a.initializeRoutes()
}

// Run the app
func (a *App) Run(port int) error {
	loggedRouter := handlers.LoggingHandler(os.Stdout, a.Router)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Router, // Pass our instance of gorilla/m
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
}
