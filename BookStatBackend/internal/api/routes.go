package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"database/sql"
)

func NewApiRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
    r.Use(middleware.Logger)

	// GET endpoints
	r.Get("/shelves",GetShelvesHandler(db))

	r.Get("/shelves/{id}",GetShelfByIdHandler(db))
	r.Get("/books",GetBooksHandler(db))
	r.Get("/books/{id}",GetBookByIdHandler(db))

	//POST endpoints
	r.Post("/shelves",CreateShelveHandler(db))
	r.Post("/books",CreateBookHandler(db))

	//PUT endpoints
	r.Put("/shelves/{id}",UpdateShelveByIdHandler(db))
	r.Put("/books/{id}",UpdateBookByIdHandler(db))

	//DELETE endpoints
	r.Put("/shelves/{id}",RemoveShelfByIdHandler(db))
	r.Put("/books/{id}",RemoveBookByIdHandler(db))

	return r
}
