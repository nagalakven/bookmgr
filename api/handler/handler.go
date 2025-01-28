package handler

import "bookmgr/pkg/service"

// Handles HTTP requests for books
type BookHandler struct {
	bookService service.BookService
}

// Initializes a new BookHandler
func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}
