package tapper

import "github.com/gdamore/tcell/v2"

type Drawable interface {
	Draw(screen tcell.Screen, layer int)
	GetDimension() (int, int)
}

type CellBuffer 