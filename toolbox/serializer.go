package toolbox

import (
	"encoding/json"
	"fmt"
)

// Serialize the message to JSON
func SerializeMessage(message interface{}) ([]byte, error) {
	serialized, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize message: %w", err)
	}
	return serialized, nil
}
