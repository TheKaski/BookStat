package api

import ( 
	"database/sql"
	"net/http"
//	"BookStatBackend/internal/models"
)

func GetShelvesHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TOdo call db  and handle response

	}
}


func GetShelfByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TOdo call db  and handle response

	}
}

func CreateShelveHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TOdo call db  and handle response

	}
}

func UpdateShelveByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TOdo call db  and handle response

	}
}

func RemoveShelfByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TOdo call db  and handle response

	}
}

