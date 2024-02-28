package handlers

import (
	"GROUPIETRACKER/pkg/render"
	"net/http"
	"GROUPIETRACKER/pkg/api"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	artists := api.GetArtists()
	render.RenderTemplate(w, "home.html", artists)
}


func ArtistDetails(w http.ResponseWriter, r *http.Request) {
	// Retrieve the artist ID from the query parameters
	artistID := r.FormValue("id")

	// Fetch the artist details from the database or your data source
	artist := api.FetchArtistDetails(artistID)

	render.RenderTemplate(w, "details.html", artist)

}
