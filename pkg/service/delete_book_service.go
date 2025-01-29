package service

import (
	"context"
	"errors"
	"fmt"

	"bookmgr/pkg/logger"
)

// Implements the DeleteBook method of the BookService interface
func (s *BookServiceImpl) DeleteBook(ctx context.Context, id string) error {

	logMsg := fmt.Sprintf("Delete Book service: id=%s\n", id)
	logger.LogMessage(logger.INFO, logMsg, false)

	// Validation
	if id == "" {
		errMsg := "validation failed: book id is required for deletion"
		logger.LogMessage(logger.ERROR, errMsg, true)
		return errors.New(errMsg)
	}

	err := s.repo.DeleteBook(id)
	if err != nil {
		errMsg := "Error deleting book: " + err.Error()
		logger.LogMessage(logger.ERROR, errMsg, true)
		return err
	}

	logger.LogMessage(logger.INFO, "Successfully deleted book with ID: "+id, false)
	return nil
}
