package db

import (
	"database/sql"
)

func InitSchema( db *sql.DB) error {

	queries := []string{
		`CREATE TABLE IF NOT EXISTS shelves (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			shelfId INTEGER NOT NULL,
			isbn TEXT NOT NULL,
			title TEXT NOT NULL,
			author TEXT,
			language TEXT,
			numberOfPages NUMBER,
			format TEXT,
			coverType TEXT,
			coverImageUrl TEXT,
			weight INTEGER NUMBER,
			FOREIGN KEY (shelfId) REFERENCES shelves(id)
		);`,
	}
	// Execute queries:
	for _, q := range queries {
		_, err := db.Exec(q)

		//If there hapens any error at any point return the err
		if err != nil {
			return err
		}
	}

	return nil
}
