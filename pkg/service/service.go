package service

import (
	"bookmgr/pkg/entity"
	"bookmgr/pkg/repository"
	"context"
)

// Interface for book services
type BookService interface {
	AddBook(ctx context.Context, book entity.Book) (*entity.Book, error)
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, id string, updatedBook entity.Book) (*entity.Book, error)
	ListBooks(ctx context.Context) ([]entity.Book, error)
}

// Implements the BookService interface
type BookServiceImpl struct {
	repo repository.Repository
}

// Initializes a BookServiceImpl
func NewBookService(repo repository.Repository) *BookServiceImpl {
	return &BookServiceImpl{repo: repo}
}
