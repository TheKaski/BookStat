package models

type Shelf struct {
	Id 	  uint `json:"id"`
	Name  string `json:"name"`
	Books []BookMetadata `json:"books,omitempty"` 
}

type ShelfMetadata struct {
	Id 	  uint `json:"id"`
	Name  string `json:"name"`
}
