package api

import (
	"encoding/json"
	"net/http"
	"main.go/utils"
	"main.go/models"
	"main.go/db"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var authReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&authReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user exists in the database.
	user, err := db.GetUserByEmail(authReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the user exists, generate a JWT token and return it.
	if user.Password == authReq.Password {
		tokenString, err := utils.GenerateJWTToken(user.ID.Hex())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the JWT token in the response header.
		w.Header().Set("Authorization", "Bearer "+tokenString)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user in the database.
	err = db.UpdateUser(userID, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
