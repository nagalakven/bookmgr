package service

import (
	"bookmgr/pkg/entity"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) AddBook(book entity.Book) (*entity.Book, error) {
	args := m.Called(book)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Book), args.Error(1)
	}
	return nil, args.Error(1)
}

// Empty stubs for other interface functions
func (m *MockRepository) DeleteBook(id string) error {
	return nil
}

func (m *MockRepository) UpdateBook(id string, updatedBook entity.Book) (*entity.Book, error) {
	return nil, nil
}

func (m *MockRepository) ListBooks() ([]entity.Book, error) {
	return nil, nil
}

// Test cases
func TestAddBook_Success(t *testing.T) {

	mockRepo := new(MockRepository)
	bookService := NewBookService(mockRepo)

	book := entity.Book{
		Title:         "Book Management System",
		Author:        "Naga",
		PublishedDate: "1990-04-10",
	}

	mockRepo.On("AddBook", mock.Anything).Return(&book, nil)

	ctx := context.Background()
	addedBook, err := bookService.AddBook(ctx, book)

	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, addedBook, "Added book should not be nil")
	assert.Equal(t, book.Title, addedBook.Title, "Book title should match")
	assert.Equal(t, book.Author, addedBook.Author, "Book author should match")
	mockRepo.AssertExpectations(t)
}

func TestAddBook_ValidationError(t *testing.T) {

	mockRepo := new(MockRepository)
	bookService := NewBookService(mockRepo)

	book := entity.Book{
		Title:         "",
		Author:        "Naga",
		PublishedDate: "2023-04-12",
	}

	ctx := context.Background()
	addedBook, err := bookService.AddBook(ctx, book)

	assert.Error(t, err, "Expected validation error")
	assert.Nil(t, addedBook, "Added book should be nil")
	assert.Equal(t, "validation failed: title, author and published date are required", err.Error())
}

func TestAddBook_RepositoryError(t *testing.T) {

	mockRepo := new(MockRepository)
	bookService := NewBookService(mockRepo)

	book := entity.Book{
		Title:         "Book Management System",
		Author:        "Naga",
		PublishedDate: "2025-01-01",
	}

	mockRepo.On("AddBook", mock.Anything).Return(nil, errors.New("database error"))

	ctx := context.Background()
	addedBook, err := bookService.AddBook(ctx, book)

	assert.Error(t, err, "Expected repository error")
	assert.Nil(t, addedBook, "Added book should be nil")
	assert.Contains(t, "Error adding book: database error", err.Error())
	mockRepo.AssertExpectations(t)
}
