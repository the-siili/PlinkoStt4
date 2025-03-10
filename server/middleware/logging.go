package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging Middleware
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler
		log.Printf("Completed in %v", time.Since(start))
	})
}
