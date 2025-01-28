package main

import (
	"fmt"
	"log"
	"net/http"

	"bookmgr/pkg/routes"
)

func main() {
	// Initialize the routes
	router := routes.SetupRoutes()

	// Start the HTTP server
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
