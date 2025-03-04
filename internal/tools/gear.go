// internal/tools/gear.go
package tools

import "os"

// Gear is a composite interface that includes multiple tool sets.
type Gear interface {
	Debugger
	FileTool
}

// gear is the concrete implementation of Gear.
type gear struct {
	debugger Debugger
	fileTool FileTool
}

// NewGear initializes a new Gear that aggregates multiple tool sets.
func NewGear() Gear {
	return &gear{
		debugger: NewDebugger(),
		fileTool: NewFileTool(),
	}
}

// Debugger methods delegate to the embedded debugger.
func (g *gear) Print(data interface{}, keys ...string) {
	g.debugger.Print(data, keys...)
}

func (g *gear) TrackRuntime(funcName string) func() {
	return g.debugger.TrackRuntime(funcName)
}

func (g *gear) TrackRoutines() {
	g.debugger.TrackRoutines()
}

// FileTool methods delegate to the embedded fileTool.
func (g *gear) ReadFile(path string) ([]byte, error) {
	return g.fileTool.ReadFile(path)
}

func (g *gear) WriteFile(path string, data []byte, perm os.FileMode) error {
	return g.fileTool.WriteFile(path, data, perm)
}
