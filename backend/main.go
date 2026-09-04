package main

import (
	"careerpilot/db"
	"careerpilot/handlers"
	"careerpilot/logger"
	"careerpilot/middleware"
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
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	logger.Init()
	db.Connect()

	http.HandleFunc("/api/health", middleware.LogRequests(withCORS(healthHandler)))
	http.HandleFunc("/api/signup", middleware.LogRequests(withCORS(middleware.RateLimit(handlers.Signup))))
	http.HandleFunc("/api/login", middleware.LogRequests(withCORS(middleware.RateLimit(handlers.Login))))
	http.HandleFunc("/api/me", middleware.LogRequests(withCORS(middleware.RequireAuth(handlers.Me))))
	http.HandleFunc("/api/upload-cv", middleware.LogRequests(withCORS(middleware.RequireAuth(handlers.UploadCV))))
	http.HandleFunc("/api/cvs", middleware.LogRequests(withCORS(middleware.RequireAuth(handlers.ListCVs))))
	http.HandleFunc("/api/cvs/", middleware.LogRequests(withCORS(middleware.RequireAuth(handlers.DeleteCV))))

	logger.Log.Info("backend starting", "addr", "http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Log.Error("server failed", "error", err)
		panic(err)
	}
}
