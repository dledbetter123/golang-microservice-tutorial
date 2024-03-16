package main

import (
	"fmt"
	"golang-microservice-tutorial/internal/userService/handler"
	"net/http"
)

// corsMiddleware adds CORS headers to api handler responses
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
	http.HandleFunc("/users", corsMiddleware(handler.GetUsers))

	fmt.Println("UserService running on port :8080")
	http.ListenAndServe(":8080", nil)
}
