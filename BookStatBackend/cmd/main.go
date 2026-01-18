package main

import (
    "log"
	"net/http"

	"BookStatBackend/internal/api"
	"BookStatBackend/internal/db"
)

func main() {

	// Connect to a db
	database, err := db.Connect("./data/data.db")
	if err != nil {
		log.Fatal(err);
	}
	defer database.Close()

	//Setup the schema if not already setup:
	if err := db.InitSchema(database); err != nil {
		log.Fatal(err)
	}

	r := api.NewApiRouter(database);
  
    http.ListenAndServe(":8080", r)
}

