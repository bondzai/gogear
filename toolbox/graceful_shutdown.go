package toolbox

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulShutdown handles the graceful shutdown of the server and multiple io.Closer resources.
func GracefulShutdown(server *http.Server, closers []io.Closer, timeout time.Duration) error {
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulStop
	log.Println("Shutdown signal received")

	// Context with timeout for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}
	log.Println("Server gracefully stopped")

	// Close all provided io.Closer resources
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			log.Printf("Error closing resource: %v", err)
		} else {
			log.Println("Resource gracefully closed")
		}
	}

	return nil
}

// Rest of your main function (database connections, message queue setup, HTTP handlers, etc.)
