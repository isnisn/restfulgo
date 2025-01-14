package handlers

import (
	"net/http"
	"rest_api_go/auth"
	"rest_api_go/services"

	"github.com/gorilla/mux"
)

var userService *services.UserService

// Inject the user service into the handlers
func SetUserService(service *services.UserService) {
	userService = service
}

// @Param username formData string true "Username"
// @Success 200 {object} string "Token"
// @Failure 400 {object} string "Invalid credentials"
// @Router /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	isValid, err := userService.Authenticate(username, password)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if isValid {
		token, err := auth.GenerateJWT(username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"token":"` + token + `"}`))

		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	isCreated, err := userService.CreateUser(username, password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if isCreated {
		w.WriteHeader(http.StatusCreated)
		return
	}
	http.Error(w, "Failed to create user", http.StatusInternalServerError)
}

func RemoveUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	// If the DeleteUser function expects a string ID
	success, err := userService.DeleteUser(username)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	if !success {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
