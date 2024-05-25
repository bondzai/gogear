package toolbox

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulShutdown handles graceful shutdown of the application.
// It receives a context, a cleanup function, and a timeout duration for the shutdown.
func GracefulShutdown(ctx context.Context, cleanup func() error, timeout time.Duration) error {
	// Create a channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal
	select {
	case <-stop:
		fmt.Println("Shutdown signal received")

		// Create a new context with timeout for the cleanup function
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		// Run the cleanup function
		done := make(chan error, 1)
		go func() {
			done <- cleanup()
		}()

		// Wait for the cleanup to complete or context timeout
		select {
		case err := <-done:
			if err != nil {
				return fmt.Errorf("cleanup error: %w", err)
			}
			fmt.Println("Cleanup completed successfully")
			return nil
		case <-ctx.Done():
			return fmt.Errorf("shutdown timeout: %w", ctx.Err())
		}
	case <-ctx.Done():
		return fmt.Errorf("context canceled: %w", ctx.Err())
	}
}
