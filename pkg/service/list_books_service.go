package service

import (
	"context"

	"bookmgr/pkg/entity"
	"bookmgr/pkg/logger"
)

// Implements the GetBooks method of the BookService interface
func (s *BookServiceImpl) ListBooks(ctx context.Context) ([]entity.Book, error) {

	booksList, err := s.repo.ListBooks()
	if err != nil {
		errMsg := "Error fetching all books: " + err.Error()
		logger.LogMessage(logger.ERROR, errMsg, true)
		return nil, err
	}

	logger.LogMessage(logger.INFO, "Successfully fetched all exsiting books", false)
	return booksList, nil
}
