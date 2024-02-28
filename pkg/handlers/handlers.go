package handlers

import (
	"GROUPIETRACKER/pkg/api"
	"GROUPIETRACKER/pkg/render"
	"net/http"
	"strings"
)

// LocationData struct holds information about a location and its date
type LocationData struct {
	Location string
	Date     string
}

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	artists := api.GetArtists()
	render.RenderTemplate(w, "home.html", artists)
}

func ArtistDetails(w http.ResponseWriter, r *http.Request) {
	// Retrieve the artist ID from the query parameters
	artistID := r.FormValue("id")

	// Fetch the artist details from the database or your data source
	artist, err := api.FetchArtistDetails(artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch relation data
	relationData, err := api.FetchRelationData("https://groupietrackers.herokuapp.com/api/relation/" + artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a struct to hold the data needed by the template
	type TemplateData struct {
		Name          string
		Image         string
		Members       []string
		CreationDate  int
		FirstAlbum    string
		LocationsData []LocationData
	}

	// Populate TemplateData struct
	var templateData TemplateData
	templateData.Name = artist.Name
	templateData.Image = artist.Image
	templateData.Members = artist.Members
	templateData.CreationDate = artist.CreationDate
	templateData.FirstAlbum = artist.FirstAlbum

	// Populate LocationsData slice
	for location, dates := range relationData.DatesLocations {
		location = strings.ReplaceAll(strings.ToLower(location), "_", " ")
		templateData.LocationsData = append(templateData.LocationsData, LocationData{
			Location: strings.Title(strings.ReplaceAll(location, "-", " ")),
			Date:     strings.Join(dates, ", "),
		})
	}

	// Render the template with the data
	render.RenderTemplate(w, "details.html", templateData)
}
