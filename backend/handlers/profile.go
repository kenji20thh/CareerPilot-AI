package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

const maxCVSize = 5 << 20 // 5MB

func UploadCV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(middleware.userIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Could Not Parse Form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("cv")
	if err != nil {
		http.Error(w, "CV file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// size check now that we actually have `header`
	if header.Size > maxCVSize {
		http.Error(w, "File too large (max 5MB)", http.StatusBadRequest)
		return // <-- this was missing
	}

	if err := os.MkdirAll("uploads", 0755); err != nil {
		http.Error(w, "Could not create uploads directory", http.StatusInternalServerError)
		return
	}
	filename := filepath.Base(header.Filename)
	path := filepath.Join("uploads", filename)

	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = dst.ReadFrom(file)
	if err != nil {
		http.Error(w, "Could not write file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"filename": filename,
		"message":  "CV uploaded successfully",
	})
}
