package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User struct represents a user entity in the API.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Users slice stores a list of user entities.
var Users []User

func main() {
	// Initialize the router.
	r := mux.NewRouter()

	// Define the API endpoints.
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUserById).Methods("GET")

	// Start the server.
	log.Fatal(http.ListenAndServe(":8080", r))
}

// getUsers returns a list of all users.
func getUsers(w http.ResponseWriter, r *http.Request) {
	// Write the users list as JSON to the response body.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

// createUser creates a new user entity.
func createUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body as a JSON-encoded User object.
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the ID of the new user as the next integer after the highest existing ID.
	nextID := 1
	for _, u := range Users {
		if u.ID >= nextID {
			nextID = u.ID + 1
		}
	}
	user.ID = nextID

	// Add the new user to the users list.
	Users = append(Users, user)

	// Write the new user as JSON to the response body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// getUserById returns a user entity by ID.
func getUserById(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL path parameter.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the user entity with the given ID.
	var user User
	found := false
	for _, u := range Users {
		if u.ID == id {
			user = u
			found = true
			break
		}
	}

	// If no user entity is found, return a 404 Not Found response.
	if !found {
		http.NotFound(w, r)
		return
	}

	// Write the user entity as JSON to the response body.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
