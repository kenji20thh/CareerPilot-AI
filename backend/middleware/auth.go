package middleware

import "net/http"

type contextKey string

const userIDKey contextKey

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
