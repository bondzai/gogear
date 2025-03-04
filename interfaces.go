// gogear/gogear.go
package gogear

import (
	"github.com/bondzai/gogear/internal/tools"
)

// gears is a global default instance of Gear.
var gears tools.Gear = tools.NewGear()

// Debugging utilities.
func Print(data interface{}, keys ...string) {
	gears.Print(data, keys...)
}

func TrackRuntime(funcName string) func() {
	return gears.TrackRuntime(funcName)
}

func TrackRoutines() {
	gears.TrackRoutines()
}

// Other utilities.
func Beta() {
	gears.Beta()
}
