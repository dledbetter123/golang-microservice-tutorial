package handler

import (
	"encoding/json"
	"net/http"
)

// Track represents a music track
type Track struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// GetTracks returns a list of tracks
func GetTracks(w http.ResponseWriter, r *http.Request) {
	tracks := []Track{
		{ID: "1", Title: "Track One", Artist: "Artist One"},
		{ID: "2", Title: "Track Two", Artist: "Artist Two"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}
