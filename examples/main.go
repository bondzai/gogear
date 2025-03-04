package main

import (
	"github.com/bondzai/gogear"
)

func main() {
	// Use debugger utilities.
	data := map[string]interface{}{
		"app":   "GoGear",
		"ver":   "1.0.0",
		"notes": "composite gear example",
	}
	gogear.Print(data)

	defer gogear.TrackRuntime("main")()
	gogear.TrackRoutines()

	// Use other utilities.
}
