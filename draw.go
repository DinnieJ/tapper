package tapper

type Drawable interface {
	Draw(x1 int, y1 int, x2 int, y2 int, content string)
}
