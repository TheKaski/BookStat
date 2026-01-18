package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"database/sql"
)

func NewApiRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
    r.Use(middleware.Logger)

	// Global CORS middleware
	r.Use(cors.Handler(cors.Options{ 
		AllowedOrigins: []string{"*"},
		// or specify "http://localhost:5173" 
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	// GET endpoints
	r.Get("/shelves",GetShelvesHandler(db))
	r.Get("/shelves/{id}",GetShelfByIdHandler(db))
	r.Get("/books",GetBooksHandler(db))
	r.Get("/books/{id}",GetBookByIdHandler(db))
	r.Get("/books/unassigned", GetUnassignedBooksHandler(db))

	//POST endpoints
	r.Post("/shelves",CreateShelveHandler(db))
	r.Post("/books",CreateBookHandler(db))

	//PUT endpoints
	r.Put("/shelves/{id}",UpdateShelveByIdHandler(db))
	r.Put("/books/{id}",UpdateBookByIdHandler(db))

	//DELETE endpoints
	r.Delete("/shelves/{id}",RemoveShelfByIdHandler(db))
	r.Delete("/books/{id}",RemoveBookByIdHandler(db))

	return r
}
