// trackperformance.go
package toolbox

import (
	"log"
	"time"
)

// TrackPerformance returns a function that when called, logs the time
// taken since TrackPerformance was called.
func TrackPerformance(funcName string) func() {
	timeStart := time.Now()
	return func() {
		timeEnd := time.Now()
		log.Printf("%s took %v to run.\n", funcName, timeEnd.Sub(timeStart))
	}
}
