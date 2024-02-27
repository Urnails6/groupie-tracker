package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
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

func GetArtists() []Artist {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error fetching data from API: %v\n", err)
		return nil
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}

	var artists []Artist
	err = json.Unmarshal(responseData, &artists)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return nil
	}

	return artists
}