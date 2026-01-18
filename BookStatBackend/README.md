# BookStatBackend - Documentation

## Introduction

BookStatBackend is a minimal GO-based REST API for managing personal book collection.

The API Exposes two core entities:

- **Shelves** - logical groups of books
- **books** - Each book can belong to a shelve. However book can also be unassigned to any shelve for easier rearrangement.

The backend uses a lightweight SQLite database stored locally in data/data.db file. Routing is implemented using the fully Go compatitible **chi** router.

## Requirements

- Go 1.21+
- SQLite (no external server required)

## Running the backend

Navigate to the project root:

`cd BookStatBackend/
go run cmd/main.go`

On the first run this will automatically create:
`BookStatBackend/data/data.db`
with the defined predefined schema if it does not yet exist.
This database starts empty and ready for use.

## Optional: Insert Mock Data

To test this API you can insert mockdata into the DB by simply running the standalone scripts/mockData.go file by running the command after the above script has created the db file for you:

`go run scripts/mockData.go `

This script inserts example shelves and books for testing. Though this data does not contain example data for all the fields. You can edit the script to insert more precise data if you like

## API Endpoints

### Shelves

| Method     | Endpoint        | Description                                     |
| ---------- | --------------- | ----------------------------------------------- |
| **GET**    | `/shelves`      | Get all shelves (metadata only)                 |
| **GET**    | `/shelves/{id}` | Get a single shelf with list of books(metadata) |
| **POST**   | `/shelves`      | Create a new shelf                              |
| **PUT**    | `/shelves/{id}` | Update a existing shelf                         |
| **DELETE** | `/shelves/{id}` | Delete a existing shelf                         |

### Books

| Method     | Endpoint            | Description                                           |
| ---------- | ------------------- | ----------------------------------------------------- |
| **GET**    | `/books`            | Get all books (metadata only)                         |
| **GET**    | `/books/{id}`       | Get a single book (full detail)                       |
| **GET**    | `/books/unassigned` | Get a list of books currently not placed in any shelf |
| **POST**   | `/books`            | Create a new book                                     |
| **PUT**    | `/books/{id}`       | Update a existing book                                |
| **DELETE** | `/books/{id}`       | Delete a existing book                                |

---
