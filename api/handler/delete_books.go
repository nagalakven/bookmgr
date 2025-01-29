package handler

import (
	"bookmgr/api/response"
	"net/http"

	"github.com/gorilla/mux"
)

// Handles the HTTP DELETE request to remove a book by ID
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := h.bookService.DeleteBook(r.Context(), id)
	if err != nil {
		response.WriteErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	responseMsg := map[string]string{"message": "Book deleted successfully!"}
	response.WriteSuccessResponse(w, response.Response{Data: responseMsg}, http.StatusOK)
}
