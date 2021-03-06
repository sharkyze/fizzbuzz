package http

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger.Printf("%s %s took=%s", r.Method, r.RequestURI, time.Since(start).String())
	})
}
