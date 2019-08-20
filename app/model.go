// model.go

package app

import (
	"database/sql"
)

type widget struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func (w *widget) getWidget(db *sql.DB) error {
	return db.QueryRow("SELECT name, price, description FROM widgets WHERE id=$1",
		w.ID).Scan(&w.Name, &w.Price, &w.Description)
}

func listWidgets(db *sql.DB, start, count int) ([]widget, error) {
	rows, err := db.Query(
		"SELECT id, name, price FROM widgets LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	widgets := []widget{}

	for rows.Next() {
		var w widget
		if err := rows.Scan(&w.ID, &w.Name, &w.Price); err != nil {
			return nil, err
		}
		widgets = append(widgets, w)
	}

	return widgets, nil
}
