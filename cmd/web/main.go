package main

import (
	"fmt"
	"net/http"

	"GROUPIETRACKER/pkg/handlers"
)

const portNumber = ":8090"

func main() {
	// Serve static files from the "static" directory
	FileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	// Handle routes
	http.HandleFunc("/", handlers.Home)
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	http.HandleFunc("/details", handlers.ArtistDetails)

	// Start the server
	_ = http.ListenAndServe(portNumber, nil)
}
