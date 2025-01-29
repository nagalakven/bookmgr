package main

import (
	"bookmgr/api/handler"
	"bookmgr/pkg/config"
	"bookmgr/pkg/constant"
	"bookmgr/pkg/logger"
	"bookmgr/pkg/repository"
	"bookmgr/pkg/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func initHTTPServer(r *mux.Router) *http.Server {
	srv := &http.Server{Handler: r}

	srv.Addr = fmt.Sprintf(":%d", config.AppConfig.HTTP.Port)
	srv.ReadTimeout = time.Duration(config.AppConfig.HTTP.ReadTimeout) * time.Second
	srv.WriteTimeout = time.Duration(config.AppConfig.HTTP.WriteTimeout) * time.Second
	srv.MaxHeaderBytes = config.AppConfig.HTTP.MaxHeaderBytes

	logMsg := fmt.Sprintf("HTTP server params - addr[%s] readTimeout[%d] writeTimeout[%d] maxHdrBytes[%d]\n",
		srv.Addr, srv.ReadTimeout, srv.WriteTimeout, srv.MaxHeaderBytes)
	logger.LogMessage(logger.INFO, logMsg, false)

	return srv
}

// Initializes and returns the router with all routes
func setupRoutes(bookHandler *handler.BookHandler) *mux.Router {
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/api/books", bookHandler.GetBooks).Methods("GET")
	router.HandleFunc("/api/books", bookHandler.AddBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	return router
}

func main() {
	// Initialize repository
	repo := repository.NewBookRepository()

	// Initialize service
	bookService := service.NewBookService(repo)

	// Initial handler
	bookHandler := handler.NewBookHandler(bookService)

	// Initialize the routes
	router := setupRoutes(bookHandler)

	// Load configurations
	config.LoadConfig(constant.CFG_PATH)

	// Start HTTP server
	logger.LogMessage(logger.INFO, "Starting Bookmgr service...", true)
	if err := (initHTTPServer(router).ListenAndServe()); err != nil {
		logger.LogMessage(logger.ERROR, "Could not start server: "+err.Error(), true)
	}
}
