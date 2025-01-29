package repository

import (
	"bookmgr/pkg/entity"
	"bookmgr/pkg/logger"
	"errors"
	"fmt"
	"sync"
)

// TODO: In-memory store for demo (this has to be replaced with DB interaction)

// Interface for data storage operations
type Repository interface {
	AddBook(book entity.Book) (*entity.Book, error)
	DeleteBook(id string) error
	UpdateBook(id string, updatedBook entity.Book) (*entity.Book, error)
	ListBooks() ([]entity.Book, error)
}

// In-memory implementation of Repository
type BookRepository struct {
	booksDB map[string]entity.Book
	mu      sync.RWMutex
}

// Iinitializes a new BookRepository
func NewBookRepository() *BookRepository {
	return &BookRepository{
		booksDB: make(map[string]entity.Book),
	}
}

// Add a new book to the in-memory storage
func (r *BookRepository) AddBook(book entity.Book) (*entity.Book, error) {

	//Lock DB before add
	r.mu.Lock()
	defer r.mu.Unlock()

	//validate if book exists
	if _, exists := r.booksDB[book.ID]; exists {
		errMsg := "Book " + book.ID + " already exists."
		logger.LogMessage(logger.ERROR, errMsg, true)
		return nil, errors.New(errMsg)
	}

	//add book to DB
	r.booksDB[book.ID] = book

	//return added content from DB
	val := r.booksDB[book.ID]

	logMsg := fmt.Sprintf("Successfully added book = %+v\n", val)
	logger.LogMessage(logger.INFO, logMsg, false)

	return &val, nil
}

func (r *BookRepository) DeleteBook(id string) error {

	//Lock DB before delete
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.booksDB[id]; exists {

		//delete book from DB
		delete(r.booksDB, id)
		logger.LogMessage(logger.INFO, "Successfully deleted book with ID: "+id, false)
		return nil
	} else {
		//Book not found, error
		errMsg := "Book [" + id + "] not found."
		logger.LogMessage(logger.ERROR, errMsg, true)
		return errors.New(errMsg)
	}
}

func (r *BookRepository) UpdateBook(id string, updatedBook entity.Book) (*entity.Book, error) {

	//Lock DB before update
	r.mu.Lock()
	defer r.mu.Unlock()

	//validate if book exists
	if currBook, exists := r.booksDB[id]; exists {

		//retain auto-generated fields
		updatedBook.ID = currBook.ID
		updatedBook.CreatedAt = currBook.CreatedAt

		//update DB
		r.booksDB[id] = updatedBook

		//return updated val from DB
		val := r.booksDB[id]

		logMsg := fmt.Sprintf("Successfully updated book = %+v\n", val)
		logger.LogMessage(logger.INFO, logMsg, false)

		return &val, nil
	} else {
		//book not found, error
		errMsg := "Book [" + id + "] not found."
		logger.LogMessage(logger.ERROR, errMsg, true)
		return nil, errors.New(errMsg)
	}
}

func (r *BookRepository) ListBooks() ([]entity.Book, error) {

	// Lock DB before read
	r.mu.RLock()
	defer r.mu.RUnlock()

	var booksList []entity.Book

	for _, book := range r.booksDB {
		booksList = append(booksList, book)
	}

	return booksList, nil
}
