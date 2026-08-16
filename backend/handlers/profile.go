package handlers

import (
	"net/http"
	"os"
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
}
