package main

import (
	"fmt"
	"net/http"

	"GROUPIETRACKER/pkg/handlers"
)

const portNumber = ":8090"

func main() {
	http.HandleFunc("/", handlers.Home)
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	http.HandleFunc("/details", handlers.ArtistDetails)
	_ = http.ListenAndServe(portNumber, nil)
}
