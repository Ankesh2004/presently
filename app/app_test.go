package app_test

import (
	"os"
	"testing"

	"github.com/dvonlehman/starter-api/app"
)

var testApp app.App

// This is a
func TestMain(m *testing.M) {
	testApp = app.App{}

	testApp.Initialize()
	testApp.Run(9998)

	os.Exit(m.Run())
}
