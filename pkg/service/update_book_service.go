package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"bookmgr/pkg/entity"
	"bookmgr/pkg/logger"
)

// Implements the UpdateBook method of the BookService interface
func (s *BookServiceImpl) UpdateBook(ctx context.Context, id string, updatedBook entity.Book) (*entity.Book, error) {

	logMsg := fmt.Sprintf("Update Book service: id=%s, book=%+v\n", id, updatedBook)
	logger.LogMessage(logger.INFO, logMsg, false)

	// Validation
	if id == "" {
		errmsg := "validation failed: book ID is required for updation"
		logger.LogMessage(logger.ERROR, errmsg, true)
		return nil, errors.New(errmsg)
	}

	// set updated time and id
	updatedBook.UpdatedAt = time.Now()

	recvBook, err := s.repo.UpdateBook(id, updatedBook)
	if err != nil {
		errMsg := "Error updating book: " + err.Error()
		logger.LogMessage(logger.ERROR, errMsg, true)
		return nil, err
	}

	logger.LogMessage(logger.INFO, "Successfully updated book with ID: "+id, false)
	return recvBook, nil
}
