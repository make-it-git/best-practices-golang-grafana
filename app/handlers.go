package main

import (
	"math/rand"
	"net/http"
	"time"
)

func successHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(50+rand.Intn(50)) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	})
}

func errorHandler(errorPercentage float64) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		if rand.Float64() < errorPercentage {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}

func slowHandler(slowDurationMax time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := rand.Intn(int(slowDurationMax))
		time.Sleep(time.Duration(n))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("slow response"))
	})
}
