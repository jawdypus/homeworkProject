package middleware

import (
	"net/http"
	"os"
	"strings"
)

func getAllowedOrigins() []string {
	// Get origins from environment variable
	originsEnv := os.Getenv("ALLOWED_ORIGINS")
	if originsEnv == "" {
		// Default origins for development
		return []string{
			"http://127.0.0.1:3000",
			"http://127.0.0.1:8080",
		}
	}

	// Split the environment variable by comma
	return strings.Split(originsEnv, ",")
}

func CORS(next http.Handler) http.Handler {

	allowedOrigins := getAllowedOrigins()

	originsMap := make(map[string]bool)
	for _, origin := range allowedOrigins {
		originsMap[origin] = true
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != "" {
			if len(allowedOrigins) == 0 || allowedOrigins[0] == "*" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else if originsMap[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin") // Important for caching
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
