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
