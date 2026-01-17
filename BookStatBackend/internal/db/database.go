package db 

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)


func Connect(dbPath string) (*sql.DB, error) {

	//Make sure the given db folder exists:
	dir := filepath.Dir(dbPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	db, error := sql.Open("sqlite", dbPath);
	
	if error != nil {
		return nil, error
	}

	return db, nil

}
