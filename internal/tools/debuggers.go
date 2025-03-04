package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"
)

// Debugger defines the contract for our debug utilities.
type (
	Debugger interface {
		// Print prints the given data.
		Print(data interface{}, keys ...string)

		// TrackRuntime returns a function to track execution time of a given function.
		TrackRuntime(funcName string) func()

		// TrackRoutines logs the current number of goroutines.
		TrackRoutines()
	}

	debugger struct{}
)

// New returns a new instance of DebugTool.
func NewDebugger() Debugger {
	return &debugger{}
}

// Print pretty prints the given data. Optionally filters output by keys.
func (d *debugger) Print(data interface{}, keys ...string) {
	if len(keys) == 0 {
		jsonBytes, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			log.Printf("Error while marshaling data to JSON: %v", err)
			return
		}
		log.Println(string(jsonBytes))
		return
	}

	// If keys are provided, assume data is a map and filter.
	mapData, ok := data.(map[string]interface{})
	if !ok {
		log.Println("Data is not a map, cannot filter keys")
		return
	}

	filteredData := make(map[string]interface{})
	for _, key := range keys {
		if value, exists := mapData[key]; exists {
			filteredData[key] = value
		}
	}

	jsonBytes, err := json.MarshalIndent(filteredData, "", "\t")
	if err != nil {
		log.Printf("Error while marshaling filtered data to JSON: %v", err)
		return
	}
	log.Println(string(jsonBytes))
}

// TrackRuntime returns a function that logs the duration since the call.
func (d *debugger) TrackRuntime(funcName string) func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		log.Printf("%s took %v to run.", funcName, duration)
	}
}

// TrackRoutines logs the number of active goroutines.
func (d *debugger) TrackRoutines() {
	fmt.Printf("Active Goroutines: %d\n", runtime.NumGoroutine())
}
