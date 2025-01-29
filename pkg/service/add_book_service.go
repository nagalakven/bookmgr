package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"bookmgr/pkg/entity"
	"bookmgr/pkg/logger"

	"github.com/google/uuid"
)

// Business logic for adding a new book
func (s *BookServiceImpl) AddBook(ctx context.Context, book entity.Book) (*entity.Book, error) {

	logMsg := fmt.Sprintf("Add Book service: book=%+v\n", book)
	logger.LogMessage(logger.INFO, logMsg, false)

	// Validation
	if book.Title == "" || book.Author == "" || book.PublishedDate == "" {
		errmsg := "validation failed: title, author and published date are required"
		logger.LogMessage(logger.ERROR, errmsg, true)
		return nil, errors.New(errmsg)
	}

	// TODO: Validate if book with same mandatory params is added again

	// Generate UUID and timestamps
	book.ID = uuid.New().String()
	book.CreatedAt = time.Now()
	book.UpdatedAt = book.CreatedAt

	addedBook, err := s.repo.AddBook(book)
	if err != nil {
		errMsg := "Error adding book: " + err.Error()
		logger.LogMessage(logger.ERROR, errMsg, true)
		return nil, err
	}

	logger.LogMessage(logger.INFO, "Successfully added book with ID: "+book.ID, false)
	return addedBook, nil
}
