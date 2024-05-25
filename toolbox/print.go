package toolbox

import (
	"encoding/json"
	"log"
	"time"
)

// PPrint pretty prints the given data object. If keysToShow is provided,
// only those keys will be included in the printed JSON.
func PPrint(data interface{}, keysToShow ...string) {
	if len(keysToShow) == 0 {
		jsonBytes, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			log.Printf("Error while marshaling data to JSON: %v", err)
		} else {
			log.Println(string(jsonBytes))
		}
	} else {
		mapData, ok := data.(map[string]interface{})
		if !ok {
			log.Println("Data is not a map, cannot filter keys")
			return
		}

		filteredData := make(map[string]interface{})
		for _, key := range keysToShow {
			if val, exists := mapData[key]; exists {
				filteredData[key] = val
			}
		}

		jsonBytes, err := json.MarshalIndent(filteredData, "", "\t")
		if err != nil {
			log.Printf("Error while marshaling filtered data to JSON: %v", err)
		} else {
			log.Println(string(jsonBytes))
		}
	}
}

// TrackRuntime returns a function that when called, logs the time
// taken since TrackRuntime was called.
func TrackRuntime(funcName string) func() {
	timeStart := time.Now()
	return func() {
		timeEnd := time.Now()
		log.Printf("%s took %v to run.\n", funcName, timeEnd.Sub(timeStart))
	}
}
