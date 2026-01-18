package db 
import ( 
	"database/sql"
	"BookStatBackend/internal/models"
)

func GetAllShelves(database *sql.DB) ([]models.ShelfMetadata, error) {
	rows, err := database.Query(`
		SELECT id, name
		FROM shelves`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shelves []models.ShelfMetadata
	for rows.Next() {
		var s models.ShelfMetadata
		rows.Scan(&s.Id, &s.Name)
		shelves = append(shelves, s)
	}

	return shelves, nil
}


func GetShelfById(database *sql.DB, shelfId uint) (models.Shelf, error) {

	var s models.Shelf

	err := database.QueryRow(`
		SELECT id, name
		FROM shelves
		WHERE id = ?
		`, shelfId).Scan(
			&s.Id,
			&s.Name,
		)

	if err != nil {
		return models.Shelf{}, err
	}

	//Append Book Metadata:
	books, err := GetBooksMetadataByShelfId(database, shelfId)
	if err != nil {
		return models.Shelf{}, err
	}

	s.Books = books

	return s, nil

}

func CreateShelf(database *sql.DB, shelfData models.Shelf) (models.Shelf, error) {
	result, err := database.Exec(`
		INSERT INTO shelves (name)
		VALUES (?)
		`, 
			shelfData.Name,
	)

	if err != nil {
		return models.Shelf{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return models.Shelf{}, err
	}

	shelfData.Id = uint(id)

	return shelfData, nil


}

func UpdateShelfMetadataById(database *sql.DB, shelfData models.ShelfMetadata) (models.ShelfMetadata, error) {
	
	_, err := database.Exec(`
		UPDATE shelves
		SET name = ? 
		WHERE id = ?
		`,
			shelfData.Name,
			shelfData.Id,
	)

	if err != nil {
		return models.ShelfMetadata{}, err
	}

	return shelfData, nil

}

func RemoveShelfById(database *sql.DB, shelfId uint) (models.ShelfMetadata, error) {

	var shelf models.ShelfMetadata
	err := database.QueryRow(`
		SELECT id, name
		FROM shelves
		WHERE id = ?
	`, shelfId).Scan(&shelf.Id, &shelf.Name)

	if err != nil {
		return models.ShelfMetadata{}, err
	}

	// unnassign all books that belong to this shelve:
	_, err = database.Exec(`
		UPDATE books
		SET shelfId = NULL
		WHERE shelfId = ?
	`, shelfId)

	if err != nil {
		return models.ShelfMetadata{}, err
	}

	// then remove the shelve:
	_, err = database.Exec(`
		DELETE FROM shelves
		WHERE id = ?
	`, shelfId)

	if err != nil {
		return models.ShelfMetadata{}, err
	}

	return shelf, nil

}
