package handler

import (
	"bookmgr/api/response"
	"net/http"
)

// Handles the HTTP GET request to fetch all books
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {

	booksList, err := h.bookService.ListBooks(r.Context())
	if err != nil {
		response.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseMsg := map[string]interface{}{
		"message": "List of Books",
		"books":   booksList}
	response.WriteSuccessResponse(w, response.Response{Data: responseMsg}, http.StatusOK)
}
