package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := getEnv("APP_PORT", "8080")

	metrics := NewMetrics()

	mux := http.NewServeMux()

	var errorPercentage = mustParsePercentage(getEnv("ERROR_PERCENTAGE", "0.01"))
	var slowDurationMax = mustParseDuration(getEnv("SLOW_DURATION_MAX", "2000"))
	mux.Handle("/success", instrumentHandler("success", successHandler(), metrics))
	mux.Handle("/error", instrumentHandler("error", errorHandler(errorPercentage), metrics))
	mux.Handle("/slow", instrumentHandler("slow", slowHandler(slowDurationMax), metrics))
	mux.Handle("/metrics", metrics.Handler())

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  mustParseDuration(getEnv("APP_READ_TIMEOUT", "5s")),
		WriteTimeout: mustParseDuration(getEnv("APP_WRITE_TIMEOUT", "10s")),
		IdleTimeout:  mustParseDuration(getEnv("APP_IDLE_TIMEOUT", "60s")),
	}

	go func() {
		log.Printf("Server starting on :%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Server gracefully stopped")
}
