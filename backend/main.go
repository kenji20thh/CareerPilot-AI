package main

import (
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

func main() {
	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/upload-cv", handlers.UploadCV)
	println("backend running on server http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
