package main

import (
	"log"
	"net/http"
	"time"
)

// Custom ResponseWriter Wrapping
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// overriding WriteHeader() function
func (rw *responseWriter) WriteHeader(code int) {

	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rw := &responseWriter{

			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		log.Println("[REQUEST]: ", r.Method, " ", r.URL.Path)
		log.Println("Processing started.....")

		start := time.Now()

		next.ServeHTTP(rw, r)

		log.Println("[RESPONSE]: Status=", rw.statusCode, " Duration=", time.Since(start))
		if rw.statusCode == 200 || rw.statusCode == 201 || rw.statusCode == 204 {

			log.Println("Processing completed.....")
		} else {

			log.Println("Processing failure.....")
		}

	})
}
