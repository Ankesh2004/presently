package app_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dvonlehman/starter-api/app"
	"gotest.tools/assert"
)

func TestRootHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)
	assert.Equal(t, response.Code, http.StatusOK)
	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, string(body), "API is running")
}
func TestHelloHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello", nil)
	response := executeRequest(req)
	assert.Equal(t, response.Code, http.StatusOK)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	assert.Equal(t, m["greeting"], "Hello")
}

func findBookById(books []app.Book, id int) *app.Book {
	for _, b := range books {
		if b.ID == id {
			return &b
		}
	}
	return nil
}

func TestListBooksHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req)

	assert.Equal(t, response.Code, http.StatusOK)
	books := make([]app.Book, 0)
	json.Unmarshal(response.Body.Bytes(), &books)

	seedData := app.SeedData()
	for _, seedBook := range seedData {
		matchingBook := findBookById(books, seedBook.ID)
		if matchingBook != nil {
			assert.DeepEqual(t, seedBook, matchingBook)
		} else {
			t.Fatal(fmt.Printf("Did not find seed book %d in results", seedBook.ID))
		}
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	testApp.Router.ServeHTTP(rr, req)

	return rr
}
