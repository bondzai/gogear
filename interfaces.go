package gogear

type Gear interface {
	PPrint()
	TrackRuntime() func()
	TrackRoutines()
}
