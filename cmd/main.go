// cmd/main.go

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukekeveloh/language-learning-app/internal/app/"
)

func main() {
	// Initialize the application
	application := app.Initialize()

	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/signup", application.SignUpHandler).Methods("POST")
	router.HandleFunc("/signin", application.SignInHandler).Methods("POST")
	router.HandleFunc("/decks", application.GetDecksHandler).Methods("GET")
	// Add more routes as needed

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	port := ":8080"
	log.Printf("Server started at http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
