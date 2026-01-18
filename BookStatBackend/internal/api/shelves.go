package api
import ( 
	"database/sql"
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"errors"
	"BookStatBackend/internal/models"
	"BookStatBackend/internal/db"
	"strconv"
)

func GetShelvesHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		shelves, err := db.GetAllShelves(database)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		
		w.WriteHeader(http.StatusOK)
		//Return the books as json
		json.NewEncoder(w).Encode(shelves)

	}
}


func GetShelfByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")
		idUint, err := strconv.ParseUint(id, 10, 64) 
		
		if err != nil { 
			http.Error(w, "Invalid ID", http.StatusBadRequest) 
			return 
		}

		shelf, err := db.GetShelfById(database,  uint(idUint))

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "The book requested was not found", 404)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(shelf)

	}
}

func CreateShelveHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var shelfObj models.Shelf
		//Parse json body:
		err := json.NewDecoder(r.Body).Decode(&shelfObj)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)	
			return
		}

		// try to store the parsed shelfObj in db:
		result, err := db.CreateShelf(database, shelfObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)

	}
}


func UpdateShelveByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var shelfObj models.ShelfMetadata

		//Parse json body:
		err := json.NewDecoder(r.Body).Decode(&shelfObj)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)	
			return
		}

		// try to store the parsed shelfObj in db:
		result, err := db.UpdateShelfMetadataById(database, shelfObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	}
}

func RemoveShelfByIdHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")
		idUint, err := strconv.ParseUint(id, 10, 64) 
		
		if err != nil { 
			http.Error(w, "Invalid ID", http.StatusBadRequest) 
			return 
		}

		shelf, err := db.RemoveShelfById(database,  uint(idUint))

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "The shelf requested was not found", 404)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(shelf)
	}
}

