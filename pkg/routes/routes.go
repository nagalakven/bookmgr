package routes

import (
	"bookmgr/api/handler"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes and returns the router with all routes
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/api/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/api/books", handler.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", handler.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", handler.DeleteBook).Methods("DELETE")

	return router
}
