package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func instrumentHandler(endpoint string, next http.Handler, m *Metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.InFlightRequests.Inc()
		defer m.InFlightRequests.Dec()

		rec := &statusRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rec, r)

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(rec.status)

		m.HttpRequestsTotal.WithLabelValues(r.Method, endpoint, status).Inc()
		m.HttpRequestDuration.WithLabelValues(r.Method, endpoint, status).Observe(duration)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method=%s path=%s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

