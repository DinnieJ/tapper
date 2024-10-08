package tapper

import "github.com/gdamore/tcell/v2"

var DefaultBorderStyle tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

// Simple box with string content inside
type Box struct {
	X1    int
	Y1    int
	X2    int
	Y2    int
	focus bool
}

func NewBox(X1 int, Y1 int, X2 int, Y2 int, content string) *Box {
	return &Box{
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
	for r := b.Y1 + 1; r < b.Y2; r++ {
		s.SetContent(b.X1, r, tcell.RuneVLine, nil, DefaultBorderStyle)
		s.SetContent(b.X2, r, tcell.RuneVLine, nil, DefaultBorderStyle)
	}

	for c := b.X1 + 1; c < b.X2; c++ {
		s.SetContent(c, b.Y1, tcell.RuneHLine, nil, DefaultBorderStyle)
		s.SetContent(c, b.Y2, tcell.RuneHLine, nil, DefaultBorderStyle)
	}

	if b.Y1 != b.Y2 && b.X1 != b.X2 {
		// Only add corners if we need to
		s.SetContent(b.X1, b.Y1, tcell.RuneULCorner, nil, DefaultBorderStyle)
		s.SetContent(b.X2, b.Y1, tcell.RuneURCorner, nil, DefaultBorderStyle)
		s.SetContent(b.X1, b.Y2, tcell.RuneLLCorner, nil, DefaultBorderStyle)
		s.SetContent(b.X2, b.Y2, tcell.RuneLRCorner, nil, DefaultBorderStyle)
	}
	for row := b.X1 + 1; row < b.X2; row++ {
		for col := b.Y1 + 1; col < b.Y2; col++ {
			s.SetContent(row, col, ' ', nil, DefaultBorderStyle)
		}
	}

}
