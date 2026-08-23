package handlers

import (
	"encoding/json"
	"net/http"
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
}
