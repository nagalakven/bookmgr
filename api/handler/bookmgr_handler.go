package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"bookmgr/pkg/entity"

	"github.com/gorilla/mux"
)

// TODO: In-memory store for demo (this has to be replaced with DB interaction)
var books = []entity.Book{}

// Handles the GET request to fetch all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

// Handles the POST request to create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.ID = uint(len(books) + 1)
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	// Add book to the in-memory store
	books = append(books, book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// Handles the PUT request to update an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedBook entity.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, book := range books {
		if fmt.Sprintf("%d", book.ID) == params["id"] {
			updatedBook.ID = book.ID
			updatedBook.CreatedAt = book.CreatedAt
			updatedBook.UpdatedAt = time.Now()

			// Replace the old book with the updated one
			books[index] = updatedBook

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	http.Error(w, "Book not found!", http.StatusNotFound)
}

// Handles the DELETE request to remove a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if fmt.Sprintf("%d", book.ID) == params["id"] {
			books = append(books[:index], books[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found!", http.StatusNotFound)
}
