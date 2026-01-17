package api
import ( 
	"database/sql"
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"errors"
	//"BookStatBackend/internal/models"
	"BookStatBackend/internal/db"
	"strconv"
)

func GetBooksHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		books, err := db.GetAllBooksMetadata(database)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		
		w.WriteHeader(http.StatusOK)
		//Return the books as json
		json.NewEncoder(w).Encode(books)
	}
}


func GetBookByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")
		idUint, err := strconv.ParseUint(id, 10, 64) 
		
		if err != nil { 
			http.Error(w, "Invalid ID", http.StatusBadRequest) 
			return 
		}

		book, err := db.GetBookById(database,  uint(idUint))

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "The book requested was not found", 404)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
		
	}
}

func CreateBookHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TOdo call db  and handle response

	}
}

func UpdateBookByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TOdo call db  and handle response

	}
}

func RemoveBookByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TOdo call db  and handle response

	}
}

