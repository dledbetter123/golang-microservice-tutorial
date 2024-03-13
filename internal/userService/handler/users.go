package handler

import (
	"encoding/json"
	"net/http"
)

// User represents a user in the system
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetUsers returns a list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: "1", Name: "John Doe"},
		{ID: "2", Name: "Jane Doe"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
