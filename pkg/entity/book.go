package entity

import "time"

// Book represents a book entity in the system
type Book struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	PublishedDate  time.Time `json:"published_date"`
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
