package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// The App struct represents the core application object
type App struct {
	Router      *mux.Router
	DB          *mongo.Database
	MongoClient *mongo.Client
}

var mongoClientCtx context.Context

func (a *App) initializeRoutes() {
	fmt.Println("Initialize the app routes")
	// Root

	// Routes
}

// Initialize the app ---> database and routes
func (a *App) Initialize(mongoURI, dbName string) error {
	// initialise database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClientCtx = ctx // storing mongo context for performing operations later
	defer cancel()       // cancel if not connected in 10 seconds

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("MongoDB connected sucessfully!")
	a.MongoClient = client
	a.DB = client.Database(dbName)

	// initialise routes
	a.Router = mux.NewRouter() // gorilla mux router
	a.initializeRoutes()
	return nil
}

// Run the app
func (a *App) Run(port int) error {
	fmt.Printf("Run the app on port %d\n", port)
	loggedRouter := handlers.LoggingHandler(os.Stdout, a.Router)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
}
