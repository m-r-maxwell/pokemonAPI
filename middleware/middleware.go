package middleware

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

// Middleware for logging requests
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		log.Println(r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Middleware for rate limiting
func RateLimiterMiddleware(next http.Handler, limiter *rate.Limiter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}
	})
}
