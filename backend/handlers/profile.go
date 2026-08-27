package handlers

import (
	"careerpilot/db"
	"careerpilot/middleware"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const maxCVSize = 5 << 20 // 5MB

var allowedExtensions = map[string]bool{
	".pdf":  true,
	".doc":  true,
	".docx": true,
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func UploadCV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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

	if header.Size > maxCVSize {
		http.Error(w, "File too large (max 5MB)", http.StatusBadRequest)
		return
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedExtensions[ext] {
		http.Error(w, "Only, PDF, DOC and DOCX files are allowed", http.StatusBadRequest)
		return
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

	var cvID string
	err = db.Pool.QueryRow(
		context.Background(),
		"INSERT INTO cvs (user_id, filename, path) VALUES ($1, $2, $3) RETURNING id",
		userID, filename, path,
	).Scan(&cvID)

	if err != nil {
		http.Error(w, "Could not save CV record", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"cvId":     cvID,
		"filename": filename,
		"message":  "CV uploaded successfully",
	})
}
