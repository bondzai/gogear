// gogear/gogear.go
package gogear

import (
	"os"

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

// File utilities.
func FileRead(path string) ([]byte, error) {
	return gears.ReadFile(path)
}

func FileWrite(path string, data []byte, perm uint32) error {
	return gears.WriteFile(path, data, os.FileMode(perm))
}
