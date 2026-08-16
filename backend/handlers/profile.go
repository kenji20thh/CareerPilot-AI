package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func uploadCV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Could Not Parse From", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("cv")
	if err != nil {
		http.Error(w, "CV file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	os.Mkdir("uploads", 7055)

	filename := filepath.Base(header.Filename)
	path := filepath.Join("uploads", filename)

	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
	}
	defer dst.Close()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"filename": filename,
		"message":  "CV uploaded successfully",
	})
}
