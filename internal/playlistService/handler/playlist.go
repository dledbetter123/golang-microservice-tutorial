package handler

import (
	"encoding/json"
	"net/http"
)

type Playlist struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Tracks []string `json:"tracks"`
}

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	playlists := []Playlist{
		{ID: "1", Name: "Chill Vibes", Tracks: []string{"1", "2"}},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlists)
}
