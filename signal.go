package tapper

type SignalType int

const (
	SignalQuit SignalType = iota
	SignalFocus
	SignalResize
	SignalDraw
)

type Signal struct {
	Sigtype SignalType
	Data    interface{}
}
