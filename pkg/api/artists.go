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
	LocationsURL string   `json:"locations"` // Change type to string for URL
	Locations    []string `json:"-"`         // Don't include this field in JSON
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

const API = "https://groupietrackers.herokuapp.com/api/"

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

func FetchArtistDetails(artistID string) (*Artist, error) {
	// Fetch artist details from the main API
	apiURL := "https://groupietrackers.herokuapp.com/api/artists/" + artistID
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from API: %v", err)
	}
	defer response.Body.Close()

	var artist Artist
	err = json.NewDecoder(response.Body).Decode(&artist)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	// Check if LocationsURL is not empty
	if artist.LocationsURL != "" {
		locationsResponse, err := http.Get(artist.LocationsURL)
		if err != nil {
			return nil, fmt.Errorf("error fetching locations: %v", err)
		}
		defer locationsResponse.Body.Close()

		var locations struct {
			Locations []string `json:"locations"`
		}
		err = json.NewDecoder(locationsResponse.Body).Decode(&locations)
		if err != nil {
			return nil, fmt.Errorf("error decoding locations: %v", err)
		}

		// Assign the fetched locations to the artist struct
		artist.Locations = locations.Locations
	}

	return &artist, nil
}

type RelationData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func FetchRelationData(url string) (*RelationData, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching relation data: %v", err)
	}
	defer response.Body.Close()

	var relationData RelationData
	err = json.NewDecoder(response.Body).Decode(&relationData)
	if err != nil {
		return nil, fmt.Errorf("error decoding relation data: %v", err)
	}

	return &relationData, nil
}
