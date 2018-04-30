package main

import (
	"log"
	"net/http"
	"time"
)

// Provide an uniforme request log management
func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		if config.LogLevel == "error" {
			log.Printf(
				"%s\t%#v\t%s\t%s",
				req.Method,
				req.RequestURI,
				time.Since(start),
			)
		}
		if config.LogLevel == "debug" {
			log.Printf(
				"%s\t%#v\t%s\t%dB\t%s\tHeaders: %s\tPayload: %s",
				req.Method,
				req.RequestURI,
				time.Since(start),
				req.ContentLength,
				req.TransferEncoding,
				req.Header,
				req.Body,
			)
		}
		next.ServeHTTP(w, req)
	})
}
