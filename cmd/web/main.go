package main

import (
	"GROUPIETRACKER/pkg/handlers"
	"fmt"
	"net/http"

)

const portNumber = ":8090"

func main() {

	http.HandleFunc("/", handlers.Home)
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
