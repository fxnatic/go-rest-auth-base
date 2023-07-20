package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
)

//Middleware Authentication Function
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the API key from the request headers.
		apiKeyHeader := r.Header.Get("x-api-key")

		// Confirm an API key was provided.
		if apiKeyHeader == "" {
			http.Error(w, "Missing X-Api-Key", http.StatusUnauthorized)
			return
		}

		// Validate the API key.
		apiKey, ok := APIKeys.Get(apiKeyHeader)
		if !ok {
			http.Error(w, "Invalid X-Api-Key", http.StatusUnauthorized)
			return
		}

		// If the API key has a limit, check and update its usage and limit.
		if apiKey.Limit != -1 {
			// If the reset time has passed, reset the usage counter and update the reset time.
			if time.Now().After(time.Unix(0, atomic.LoadInt64(&apiKey.Reset))) {
				atomic.StoreInt64(&apiKey.Usage, 0)
				atomic.StoreInt64(&apiKey.Reset, time.Now().Add(apiKey.Duration).UnixNano())
			}

			// If the usage exceeds the limit, return 429 error.
			if atomic.LoadInt64(&apiKey.Usage) >= apiKey.Limit {
				http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
				return
			}

			// Increment the usage counter.
			atomic.AddInt64(&apiKey.Usage, 1)
		}

		next.ServeHTTP(w, r)
	})
}

// Middleware Logging Function
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Print the log to the console.
		fmt.Printf("[%s][%s][%s][%s]\n",
			color.BlueString("%s", time.Now().Format(time.RFC3339)), // Current Time
			color.GreenString("%s", r.Method),                       // Request Method
			color.YellowString(r.URL.Path),                          // Request Path
			color.HiWhiteString(r.Header.Get("x-api-key")),          // API Key
		)

		next.ServeHTTP(w, r)
	})
}
