package handlers

import (
	"careerpilot/db"
	"careerpilot/middleware"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)

type CV struct {
	ID         string    `json:"id"`
	Filename   string    `json:"filename"`
	UploadedAt time.Time `json:"uploadedAt"`
}

// ListCVs returns every CV belonging to the logged-in user.
func ListCVs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	rows, err := db.Pool.Query(
		context.Background(),
		"SELECT id, filename, uploaded_at FROM cvs WHERE user_id = $1 ORDER BY uploaded_at DESC",
		userID,
	)
	if err != nil {
		http.Error(w, "Could not fetch CVs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	cvs := []CV{}
	for rows.Next() {
		var c CV
		if err := rows.Scan(&c.ID, &c.Filename, &c.UploadedAt); err != nil {
			http.Error(w, "Error reading CV data", http.StatusInternalServerError)
			return
		}
		cvs = append(cvs, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"cvs":     cvs,
	})
}

// DeleteCV removes a CV, but only if it belongs to the logged-in user.
// Expects the route to be registered as /api/cvs/ so the ID can be
// extracted from the remainder of the path.
func DeleteCV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	cvID := strings.TrimPrefix(r.URL.Path, "/api/cvs/")
	if cvID == "" {
		http.Error(w, "CV id is required", http.StatusBadRequest)
		return
	}

	var path string
	err := db.Pool.QueryRow(
		context.Background(),
		"SELECT path FROM cvs WHERE id = $1 AND user_id = $2",
		cvID, userID,
	).Scan(&path)

	if err != nil {
		http.Error(w, "CV not found", http.StatusNotFound)
		return
	}

	_, err = db.Pool.Exec(
		context.Background(),
		"DELETE FROM cvs WHERE id = $1 AND user_id = $2",
		cvID, userID,
	)
	if err != nil {
		http.Error(w, "Could not delete CV record", http.StatusInternalServerError)
		return
	}

	// Best-effort file cleanup — the DB row is already gone, so we don't fail
	// the request if this errors, but a stale file left on disk is not ideal.
	_ = os.Remove(path)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "CV deleted successfully",
	})
}
