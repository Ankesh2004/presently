package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// RootHandler for the root url
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is running")
}

// HelloHandler route handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, map[string]string{"greeting": "Hello"})
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
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
