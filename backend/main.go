package main

import (
	"careerpilot/db"
	"careerpilot/handlers"
	"encoding/json"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	db.Connect()

	http.HandleFunc("/api/health", withCORS(healthHandler))
	http.HandleFunc("/api/upload-cv", withCORS(handlers.UploadCV))
	http.HandleFunc("/api/signup", withCORS(handlers.Signup))
	http.HandleFunc("/api/login", withCORS(handlers.Login))
	println("backend running on server http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
