package main

import (
	"log"
	"net/http"

	"myapp/config"
	"myapp/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to the database
	config.Connect()
	defer config.Close()

	// Create a new router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
