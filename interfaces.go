package gogear

import "github.com/bondzai/gogear/internal/tools"

// gears is a global default instance of DebugTool.
var gears tools.Debugger = tools.NewDebugger()

// Print is a wrapper around the DebugTool.Print method.
func Print(data interface{}, keys ...string) {
	gears.Print(data, keys...)
}

// TrackRuntime is a wrapper around the DebugTool.TrackRuntime method.
func TrackRuntime(funcName string) func() {
	return gears.TrackRuntime(funcName)
}

// TrackRoutines is a wrapper around the DebugTool.TrackRoutines method.
func TrackRoutines() {
	gears.TrackRoutines()
}
