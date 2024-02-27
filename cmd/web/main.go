package main

import (
	"GROUPIETRACKER/pkg/handlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

const portNumber = ":8090"

func main() {

	// URL of the API you want to fetch data from
	apiUrl := "https://groupietrackers.herokuapp.com/api/artists"

	// Make GET request to the API
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error fetching data from API: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Create an instance of Data struct to hold the parsed data
	var ledger []Data

	// Unmarshal the JSON response into the Data struct
	err = json.Unmarshal(responseData, &ledger)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	http.HandleFunc("/", handlers.Home)
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}

