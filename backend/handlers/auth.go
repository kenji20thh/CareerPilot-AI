package handlers

import "net/http"

type SignupRequest struct {
	Email    string `json:"email`
	Password string `json:"password"`
}

func Singup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}
