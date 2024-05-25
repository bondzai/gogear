package toolbox

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestSerializeMessage(t *testing.T) {
	t.Run("successful serialization", func(t *testing.T) {
		message := map[string]string{
			"greeting": "hello",
			"subject":  "world",
		}
		expected := `{"greeting":"hello","subject":"world"}`

		serialized, err := SerializeMessage(message)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if string(serialized) != expected {
			t.Fatalf("expected %s, got %s", expected, string(serialized))
		}
	})

	t.Run("serialization error", func(t *testing.T) {
		// json.Marshal cannot serialize channels, so this should produce an error
		message := make(chan int)

		_, err := SerializeMessage(message)
		if err == nil {
			t.Fatal("expected an error, got none")
		}

		unwrappedErr := errors.Unwrap(err)
		var jsonErr *json.UnsupportedTypeError
		if !errors.As(unwrappedErr, &jsonErr) {
			t.Fatalf("expected a json.UnsupportedTypeError, got %v", err)
		}
	})
}
