package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// RootHandler for the root url
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is running")
}

// ListBooksHandler lists out all the Books
func (a *App) ListBooksHandler(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	books, err := listBooks(a.Database, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching book: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

// GetBookHandler fetches a single book by ID
func (a *App) GetBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := getBook(a.Database, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching book: %s", err.Error()))
		return
	}
	if book == nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("Could not find Book with id %d", id))
		return
	}

	respondWithJSON(w, 200, book)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// The ErrorResponse struct describes an error JSON returned by the API
type ErrorResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	json := ErrorResponse{Error: message}
	respondWithJSON(w, code, json)
}
