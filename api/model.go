// model.go

package api

// The Book model
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishYear int    `json:"publish_year"`
}
