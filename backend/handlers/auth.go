package handlers

import (
	"careerpilot/db"
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Email    string `json:"email`
	Password string `json:"password"`
}

func Singup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and Password required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}

	var userID string
	err = db.Pool.QueryRow(
		context.Background,
		"INSERT INTO users(email, password_hash) VALUES ($1, $2) RETURNING id",
		req.Email, string(hashedPassword),
	).Scan(&userID)

	if err != nil {
		http.Error(w, "Could not create user (email might already be taken)", http.StatusConflict)
		return
	}

	token, err := generateJWT(userID)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}
}
