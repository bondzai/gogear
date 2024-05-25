package toolbox

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestGracefulShutdown(t *testing.T) {
	t.Run("successful cleanup", func(t *testing.T) {
		cleanup := func() error {
			time.Sleep(1 * time.Second) // Simulate cleanup work
			return nil
		}

		ctx := context.Background()
		err := GracefulShutdown(ctx, cleanup, 5*time.Second)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("cleanup error", func(t *testing.T) {
		cleanup := func() error {
			time.Sleep(1 * time.Second) // Simulate cleanup work
			return errors.New("cleanup failed")
		}

		ctx := context.Background()
		err := GracefulShutdown(ctx, cleanup, 5*time.Second)
		if err == nil || err.Error() != "cleanup error: cleanup failed" {
			t.Fatalf("expected cleanup error, got %v", err)
		}
	})

	t.Run("timeout", func(t *testing.T) {
		cleanup := func() error {
			time.Sleep(6 * time.Second) // Simulate long cleanup work
			return nil
		}

		ctx := context.Background()
		err := GracefulShutdown(ctx, cleanup, 3*time.Second)
		if err == nil || err.Error() != "shutdown timeout: context deadline exceeded" {
			t.Fatalf("expected timeout error, got %v", err)
		}
	})
}
