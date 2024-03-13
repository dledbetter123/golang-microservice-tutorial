package main

import (
	"fmt"
	"net/http"

	"golang-microservice-tutorial/internal/playlistService/handler"
)

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/playlists", corsMiddleware(handler.GetPlaylists))
	fmt.Println("PlaylistService running on port :8082")
	http.ListenAndServe(":8082", nil)
}
