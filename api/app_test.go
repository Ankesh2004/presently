package api_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"example.com/starter-api/api"
)

var testApp api.App

// This is a
func TestMain(m *testing.M) {
	fmt.Println("Initialize the test app")

	sqldb, err := api.CreateSqliteDatabase()
	if err != nil {
		log.Fatal(err)
	}

	app := api.App{DB: sqldb}
	if err := app.Initialize(); err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}
