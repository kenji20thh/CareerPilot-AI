package middleware

import (
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// visitor holds the rate limiter for a single IP address
type visitor struct {
	limiter *rate.Limiter
}

var (
	visitors   = make(map[string]*visitor)
	visitorsMu sync.Mutex
)

// getVisitor returns the rate limiter for a given IP, creating one if it doesn't exist yet.
// Allows 5 requests, refilling at 1 every 10 seconds (roughly 6 requests/minute sustained).
func getVisitor(ip string) *rate.Limiter {
	visitorsMu.Lock()
	defer visitorsMu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(rate.Every(10_000_000_000), 5) // 1 token per 10s, burst of 5
		visitors[ip] = &visitor{limiter: limiter}
		return limiter
	}
	return v.limiter
}

// clientIP extracts the request's originating IP, stripping the port.
func clientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// RateLimit wraps a handler, rejecting requests once the caller's IP exceeds its allowance.
func RateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := clientIP(r)
		limiter := getVisitor(ip)

		if !limiter.Allow() {
			http.Error(w, "Too many requests, please try again later", http.StatusTooManyRequests)
			return
		}

		next(w, r)
	}
}
