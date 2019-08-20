package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

	a.Router.HandleFunc("/widgets", a.ListProductsHandler).Methods("GET")
	// a.Router.HandleFunc("/widget", a.createProduct).Methods("POST")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.getProduct).Methods("GET")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	// a.Router.HandleFunc("/widget/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *App) initializeDb() {
	statement, _ := a.Database.Prepare("CREATE TABLE IF NOT EXISTS widgets (id INTEGER PRIMARY KEY, name TEXT, price NUMERIC, description TEXT)")
	statement.Exec()

	// Insert some widgets
	var testWidgets = []widget{
		widget{
			ID:          1,
			Name:        "Rubber Mallet",
			Price:       19.95,
			Description: "Your basic mallet for hitting with",
		},
		widget{
			ID:          2,
			Name:        "Allen Wrench",
			Price:       15,
			Description: "Needed for assembling IKEA furniture",
		},
	}

	// Insert the test data
	for _, w := range testWidgets {
		statement, _ := a.Database.Prepare("INSERT INTO widgets (id, name, price, description) VALUES (?, ?) WHERE NOT EXISTS (SELECT * FROM widgets WHERE id=?)")
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

	return http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
}
