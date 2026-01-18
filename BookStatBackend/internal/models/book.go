package models

type Book struct { 
	Id 			  uint	 `json:"id"`
	ShelfId 	  uint	 `json:"shelfId"`
	Isbn		  string `json:"isbn"`
	Title 		  string `json:"title"`
	Author 	 	  string `json:"author"`
	Language	  string `json:"language"`
	NumberOfPages uint	 `json:"numberOfPages"`
	Format 		  string `json:"format"`
	CoverType 	  string `json:"coverType"`
	CoverImageUrl string `json:"coverImageUrl"` 
	Weight 		  uint 	 `json:"weight"`
}


type BookMetadata struct { 
	Id 			  uint	 `json:"id"`
	ShelfId 	  uint	 `json:"shelfId"`
	Title 		  string `json:"title"`
	Author 	 	  string `json:"author"`
	CoverImageUrl string `json:"coverImageUrl"` 
}


