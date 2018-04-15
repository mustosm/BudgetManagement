package main

import (
	"log"
	"net/http"
	"time"
)

// Provide an uniforme log management
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		if config.LogLevel == "error" {
			log.Printf(
				"%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		}
		if config.LogLevel == "debug" {
			log.Printf(
				"%s\t%s\t%s\t%s\t%d\t%s:\n%s\n%s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
				r.ContentLength,
				r.TransferEncoding,
				r.Header,
				r.Body,
			)
		}
	})
}
