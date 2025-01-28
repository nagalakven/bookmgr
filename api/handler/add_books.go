package handler

import (
	"encoding/json"
	"net/http"

	"bookmgr/api/response"
	"bookmgr/pkg/entity"
)

// Handles the HTTP POST request to add a book
func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		response.WriteErrorResponse(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	createdBook, err := h.bookService.AddBook(r.Context(), book)
	if err != nil {
		response.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteSuccessResponse(w, createdBook, http.StatusCreated)
}
