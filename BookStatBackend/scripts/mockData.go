package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "modernc.org/sqlite"
)

func main() {
    db, err := sql.Open("sqlite", "./data/data.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    fmt.Println("Seeding mock data...")

    if err := seedShelves(db); err != nil {
        log.Fatal("Error seeding shelves:", err)
    }

    if err := seedBooks(db); err != nil {
        log.Fatal("Error seeding books:", err)
    }

    fmt.Println("Mock data inserted successfully.")
}

func seedShelves(db *sql.DB) error {
    shelves := []string{
        `INSERT INTO shelves (name) VALUES ('Fantasy');`,
        `INSERT INTO shelves (name) VALUES ('Science Fiction');`,
        `INSERT INTO shelves (name) VALUES ('Non-Fiction');`,
    }

    for _, q := range shelves {
        if _, err := db.Exec(q); err != nil {
            return err
        }
    }

    return nil
}

func seedBooks(db *sql.DB) error {
    books := []string{
        `INSERT INTO books (shelfId, isbn, title, author, language, numberOfPages, format, coverType, coverImageUrl, weight)
         VALUES (1, '9780261103573', 'The Lord of the Rings', 'J.R.R. Tolkien', 'English', 1216, 'Paperback', 'Softcover', '', 900);`,

        `INSERT INTO books (shelfId, isbn, title, author, language, numberOfPages, format, coverType, coverImageUrl, weight)
         VALUES (1, '9780765376671', 'Mistborn: The Final Empire', 'Brandon Sanderson', 'English', 672, 'Paperback', 'Softcover', '', 600);`,

        `INSERT INTO books (shelfId, isbn, title, author, language, numberOfPages, format, coverType, coverImageUrl, weight)
         VALUES (2, '9780553380163', 'Dune', 'Frank Herbert', 'English', 896, 'Paperback', 'Softcover', '', 700);`,

        `INSERT INTO books (shelfId, isbn, title, author, language, numberOfPages, format, coverType, coverImageUrl, weight)
         VALUES (3, '9780143127741', 'Sapiens', 'Yuval Noah Harari', 'English', 498, 'Paperback', 'Softcover', '', 500);`,
    }

    for _, q := range books {
        if _, err := db.Exec(q); err != nil {
            return err
        }
    }

    return nil
}

