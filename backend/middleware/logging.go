package middleware

import (
	"careerpilot/logger"
	"net/http"
	"time"
)

// responseRecorder wraps http.ResponseWriter so we can capture the status
// code actually written, since the standard interface doesn't expose it.
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

// LogRequests wraps a handler, logging method, path, status, and duration
// for every request that passes through it.
func LogRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next(rec, r)

		logger.Log.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.statusCode,
			"duration_ms", time.Since(start).Milliseconds(),
			"ip", r.RemoteAddr,
		)
	}
}
