package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mysterybee07/go-oauth/auth"
	"github.com/mysterybee07/go-oauth/database"
	"github.com/mysterybee07/go-oauth/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environmental variables")
	}

	// Connect to the database
	database.ConnectDB()

	// Initialize authentication (OAuth)
	auth.NewAuth()

	// Set up routes
	router := routes.Route()

	// Get the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}

	// Start the HTTP server
	log.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, router)
}
