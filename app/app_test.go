package app_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/dvonlehman/starter-api/app"
)

var testApp app.App

// This is a
func TestMain(m *testing.M) {
	fmt.Println("Initialize the test app")
	testApp = app.App{}

	testApp.Initialize()
	os.Exit(m.Run())
}
