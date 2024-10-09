package tapper

import "github.com/gdamore/tcell/v2"

type SignalType int

const (
	SignalQuit SignalType = iota
	SignalFocus
	SignalResize
	SignalDraw
	SignalCallback
	SignalSuspend
	SignalContinue
)

type Signal struct {
	Sigtype SignalType
	Data    interface{}
}

type SignalCallbackFn func(s tcell.Screen)
