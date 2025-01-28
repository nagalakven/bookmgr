package handler

import (
	"bookmgr/api/response"
	"bookmgr/pkg/entity"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handles the PUT request to update an existing book
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var updatedBook entity.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		response.WriteErrorResponse(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	recvBook, err := h.bookService.UpdateBook(r.Context(), id, updatedBook)
	if err != nil {
		response.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteSuccessResponse(w, recvBook, http.StatusOK)
}
