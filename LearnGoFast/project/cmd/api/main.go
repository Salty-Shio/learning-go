package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"go_API/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Start Logging
	log.SetReportCaller(true)

	// Initialize the router
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	// Start the server
	fmt.Println("Starting GO API service...")
	err := http.ListenAndServe("localhost:8080", r)

	// Log errors if server fails to start
	if err != nil {
		log.Error(err)		
	}

	
}