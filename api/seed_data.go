package api

// SeedData returns the set of sample books to seed the db with
func SeedData() []*Book {
	return []*Book{
		&Book{
			ID:          1,
			Title:       "Ulysses",
			Author:      "James Joyce",
			PublishYear: 1922,
		},
		&Book{
			ID:          2,
			Title:       "The Great Gatsby",
			Author:      "F Scott Fitzgerald",
			PublishYear: 1925,
		},
		&Book{
			ID:          3,
			Title:       "Moby Dick",
			Author:      "Herman Melville",
			PublishYear: 1851,
		},
		&Book{
			ID:          4,
			Title:       "War and Peace",
			Author:      "Leo Tolstoy",
			PublishYear: 1869,
		},
	}
}
