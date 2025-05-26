package app

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	delivery "quotesAPI/internal/delivery/http"
	"syscall"
	"time"
)

func (a *Application) Run() error {
	router := mux.NewRouter()
	delivery.SetupRoutes(router, a.QuoteService)

	server := &http.Server{
		Addr:    a.Port,
		Handler: router,
		// Add timeouts to avoid resource leaks
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Channel to listen for errors from server
	serverErrors := make(chan error, 1)

	// Start server
	go func() {
		log.Printf("Server started on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErrors <- err
		}
	}()

	// Channel to listen signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return err

	case sig := <-shutdown:
		log.Printf("Received %v signal, starting graceful shutdown", sig)

		// Create context with timeout for shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt graceful shutdown
		if err := server.Shutdown(ctx); err != nil {
			// force close
			log.Printf("Graceful shutdown failed: %v", err)
			if err := server.Close(); err != nil {
				return err
			}
		}
	}

	log.Println("Server stopped")
	return nil
}
