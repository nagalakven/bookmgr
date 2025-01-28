package response

import (
	"encoding/json"
	"net/http"
)

// Sends an error response
func WriteErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Sends a success response
func WriteSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if statusCode != http.StatusNoContent {
		json.NewEncoder(w).Encode(data)
	}
}
