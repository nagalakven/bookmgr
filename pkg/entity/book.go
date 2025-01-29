package entity

import (
	"context"
	"time"
)

// Book entity
type Book struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	PublishedDate  string    `json:"published_date"`
	PublishedBy    string    `json:"published_by,omitempty"`
	CoAuthors      string    `json:"co_authors,omitempty"`
	Genre          string    `json:"genre,omitempty"`
	Description    string    `json:"description,omitempty"`
	Edition        string    `json:"edition,omitempty"`
	RecommendedAge string    `json:"recommended_age,omitempty"`
	ISBN           string    `json:"isbn,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// BookService interface
type BookService interface {
	//methods
	AddBook(ctx context.Context, book Book) (Book, error)
	UpdateBook(ctx context.Context, updateBook Book) (Book, error)
	DeleteBook(ctx context.Context, id string) error
	GetBooks(ctx context.Context) ([]Book, error)
}
