package main

import (
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	// Serve static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route handlers
	http.HandleFunc("/", handlers.HomeHandler)

	// Start server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
