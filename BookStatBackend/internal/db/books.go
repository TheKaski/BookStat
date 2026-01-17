package db 
import ( 
	"database/sql"
	"BookStatBackend/internal/models"
)

func GetAllBooksMetadata(database *sql.DB) ([]models.BookMetadata, error) {
	rows, err := database.Query(`
		SELECT id, shelfId, title, author, coverImageUrl 
		FROM books`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.BookMetadata
	for rows.Next() {
		var b models.BookMetadata
		rows.Scan(&b.Id, &b.ShelfId, &b.Title, &b.Author, &b.CoverImageUrl)
		books = append(books, b)
	}

	return books, nil
}

func GetBooksMetadataByShelfId(database *sql.DB, shelfId uint ) ([]models.BookMetadata, error) {
	rows, err := database.Query(`
		SELECT id, shelfId, title, author, coverImageUrl 
		FROM books
		WHERE shelfId = ?`, shelfId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.BookMetadata
	for rows.Next() {
		var b models.BookMetadata
		rows.Scan(&b.Id, &b.ShelfId, &b.Title, &b.Author, &b.CoverImageUrl)
		books = append(books, b)
	}

	return books, nil
}



func GetBookById(database *sql.DB, bookId uint) (models.Book, error) {

	var b models.Book

	err := database.QueryRow(`
		SELECT id, shelfId, title, author, coverImageUrl 
		FROM books
		WHERE id = ?
		`, bookId).Scan(
			&b.Id,
			&b.ShelfId,
			&b.Title,
			&b.Author,
			&b.CoverImageUrl,
		)

	if err != nil {
		return models.Book{}, err
	}

	return b, nil

}

func CreateBook(database *sql.DB, bookData models.Book) (models.Book, error) {
	result, err := database.Exec(`
		INSERT INTO books (shelfId, isbn, title, author, language, numberOfPages, format, coverType, coverImageUrl, weight)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, 
			bookData.ShelfId,
			bookData.Isbn,
			bookData.Title,
			bookData.Author,
			bookData.Language,
			bookData.NumberOfPages,
			bookData.Format,
			bookData.CoverType,
			bookData.CoverImageUrl,
			bookData.Weight,
	)

	if err != nil {
		return models.Book{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return models.Book{}, err
	}

	bookData.Id = uint(id)

	return bookData, nil

}


