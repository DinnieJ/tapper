package tapper

import (
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

var DefaultBorderStyle tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorDarkGray).Background(tcell.ColorBlack)
var FocusBorderStyle tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

var FocusBorderRune = []rune{'║', '═', '╔', '╗', '╚', '╝'}

var DefaultBorderRune = []rune{tcell.RuneVLine, tcell.RuneHLine, tcell.RuneULCorner, tcell.RuneURCorner, tcell.RuneLLCorner, tcell.RuneLRCorner}

// Simple box with string content inside
type Box struct {
	ID    string
	X1    int
	Y1    int
	X2    int
	Y2    int
	focus bool
}

func NewBox(id string, X1 int, Y1 int, X2 int, Y2 int) *Box {
	if X1 > X2 {
		X1, X2 = X2, X1
	}

	if Y1 > Y2 {
		Y1, Y2 = Y2, Y1
	}
	if id == "" {
		uuid, _ := uuid.NewUUID()
		id = uuid.String()
	}
	return &Box{
		ID: id,
		X1: X1,
		Y1: Y1,
		X2: X2,
		Y2: Y2,
	}
}

func (b *Box) SetFocus(focus bool) {
	b.focus = focus
}

func (b *Box) Draw(s tcell.Screen) {
	var borderRune []rune
	if b.focus {
		borderRune = FocusBorderRune
	} else {
		borderRune = DefaultBorderRune
	}
	for r := b.Y1 + 1; r < b.Y2; r++ {
		s.SetContent(b.X1, r, borderRune[0], nil, DefaultBorderStyle)
		s.SetContent(b.X2, r, borderRune[0], nil, DefaultBorderStyle)
	}

	for c := b.X1 + 1; c < b.X2; c++ {
		s.SetContent(c, b.Y1, borderRune[1], nil, DefaultBorderStyle)
		s.SetContent(c, b.Y2, borderRune[1], nil, DefaultBorderStyle)
	}

	if b.Y1 != b.Y2 && b.X1 != b.X2 {
		// Only add corners if we need to
		s.SetContent(b.X1, b.Y1, borderRune[2], nil, DefaultBorderStyle)
		s.SetContent(b.X2, b.Y1, borderRune[3], nil, DefaultBorderStyle)
		s.SetContent(b.X1, b.Y2, borderRune[4], nil, DefaultBorderStyle)
		s.SetContent(b.X2, b.Y2, borderRune[5], nil, DefaultBorderStyle)
	}
	for row := b.X1 + 1; row < b.X2; row++ {
		for col := b.Y1 + 1; col < b.Y2; col++ {
			s.SetContent(row, col, ' ', nil, DefaultBorderStyle)
		}
	}

	// s.LockRegion(b.X1, b.Y1, b.X2-b.X1, b.Y2-b.Y1, true)

}

func (b *Box) GetDimension() (int, int) {
	return b.X2 - b.X1, b.Y2 - b.Y1
}
