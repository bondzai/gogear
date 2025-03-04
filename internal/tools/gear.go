// internal/tools/gear.go
package tools

// Gear is a composite interface that includes multiple tool sets.
type Gear interface {
	Debugger
	Beta
}

// gear is the concrete implementation of Gear.
type gear struct {
	debugger Debugger
	beta     Beta
}

// NewGear initializes a new Gear that aggregates multiple tool sets.
func NewGear() Gear {
	return &gear{
		debugger: NewDebugger(),
		beta:     NewBeta(),
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

// Beta methods delegate to the embedded beta.
func (g *gear) Beta() {
	g.beta.Beta()
}
