package clikit

import termbox "github.com/nsf/termbox-go"

// RelCoord is a coordinate relative to a Canvas's origin
type RelCoord int

// AbsCoord is a coordinate relative to the terminal's origin
type AbsCoord int

// Canvas represents an absolute rectange on the screen upon which drawing operations are applied.
type Canvas struct {
	// Absolute X
	X AbsCoord
	// Absolute Y
	Y AbsCoord
	// Width
	W int
	// Height
	H int
}

// NewCanvasFullScreen constructs a BoundingBox encompassing the entire terminal.
func NewCanvasFullScreen() Canvas {
	w, h := termbox.Size()
	return Canvas{
		X: 0,
		Y: 0,
		W: w,
		H: h,
	}
}

// relToAbs translates canvas-relative coordinates to absolute screen coordinates
func (cvs *Canvas) relToAbs(xr, yr RelCoord) (xa, ya AbsCoord) {
	return AbsCoord(int(xr) + int(cvs.X)), AbsCoord(int(yr) + int(cvs.Y))
}

// isIn determines if a pair of absolute coordinates are
func (cvs *Canvas) isIn(xa, ya AbsCoord) bool {
	return cvs.X <= xa && xa < AbsCoord(int(cvs.X)+cvs.W) && cvs.Y <= ya && ya < AbsCoord(int(cvs.Y)+cvs.H)
}

// FillBgBox fills a box with a background color, preserving the foreground contents.
func (cvs *Canvas) FillBgBox(x, y RelCoord, w, h int, bg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	for xi := x; xi < RelCoord(int(x)+w); xi++ {
		for yi := y; yi < RelCoord(int(y)+h); yi++ {
			xa, ya := cvs.relToAbs(xi, yi)
			if cvs.isIn(xa, ya) {
				curCell := cvs.cellAt(curBuf, xa, ya)
				cvs.setCell(xa, ya, curCell.Ch, curCell.Fg, bg)
			}
		}
	}
}

// FillFgBox fills a box with a rune, preserving the background.
func (cvs *Canvas) FillFgBox(x, y RelCoord, w, h int, ch rune, fg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	for xi := x; xi < RelCoord(int(x)+w); xi++ {
		for yi := y; yi < RelCoord(int(y)+h); yi++ {
			xa, ya := cvs.relToAbs(xi, yi)
			if cvs.isIn(xa, ya) {
				curCell := cvs.cellAt(curBuf, xa, ya)
				cvs.setCell(xa, ya, ch, fg, curCell.Bg)
			}
		}
	}
}

// DrawHorizLine draws a horizontal line with a given rune in the foreground,
// preserving the background.
func (cvs *Canvas) DrawHorizLine(x, y RelCoord, len int, ch rune, fg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	for xi := x; xi < RelCoord(int(x)+len); xi++ {
		xa, ya := cvs.relToAbs(xi, y)
		if cvs.isIn(xa, ya) {
			curCell := cvs.cellAt(curBuf, xa, ya)
			cvs.setCell(xa, ya, ch, fg, curCell.Bg)
		}
	}
}

// DrawCappedHorizLine draws a horizontal line with a given rune and endcaps in the foreground,
// preserving the background.
func (cvs *Canvas) DrawCappedHorizLine(x, y RelCoord, len int, left, middle, right rune, fg termbox.Attribute) {
	cvs.DrawRuneFg(x, y, left, fg)
	cvs.DrawHorizLine(x+1, y, len-2, middle, fg)
	cvs.DrawRuneFg(RelCoord(int(x)+len-1), y, right, fg)
}

// DrawHorizLineStyle draws a horizontal line with a given style in the foreground,
// preserving the background.
func (cvs *Canvas) DrawHorizLineStyle(x, y RelCoord, len int, lineStyle *LineStyle, fg termbox.Attribute) {
	cvs.DrawCappedHorizLine(x, y, len, lineStyle.LeftCap, lineStyle.Horiz, lineStyle.RightCap, fg)
}

// DrawVertLine draws a horizontal line with a given rune in the foreground,
// preserving the background.
func (cvs *Canvas) DrawVertLine(x, y RelCoord, len int, ch rune, fg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	for yi := y; yi < RelCoord(int(y)+len); yi++ {
		xa, ya := cvs.relToAbs(x, yi)
		if cvs.isIn(xa, ya) {
			curCell := cvs.cellAt(curBuf, xa, ya)
			cvs.setCell(xa, ya, ch, fg, curCell.Bg)
		}
	}
}

// DrawCappedVertLine draws a vertical line with a given rune and endcaps in the foreground,
// preserving the background.
func (cvs *Canvas) DrawCappedVertLine(x, y RelCoord, len int, top, middle, bottom rune, fg termbox.Attribute) {
	cvs.DrawRuneFg(x, y, top, fg)
	cvs.DrawVertLine(x, y+1, len-2, middle, fg)
	cvs.DrawRuneFg(x, RelCoord(int(y)+len-1), bottom, fg)
}

// DrawVertLineStyle draws a vertical line with a given style in the foreground,
// preserving the background.
func (cvs *Canvas) DrawVertLineStyle(x, y RelCoord, len int, lineStyle *LineStyle, fg termbox.Attribute) {
	cvs.DrawCappedVertLine(x, y, len, lineStyle.TopCap, lineStyle.Vert, lineStyle.BottomCap, fg)
}

// DrawRuneFg draws a rune in the foreground, preserving the background.
func (cvs *Canvas) DrawRuneFg(x, y RelCoord, ch rune, fg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	xa, ya := cvs.relToAbs(x, y)
	if cvs.isIn(xa, ya) {
		curCell := cvs.cellAt(curBuf, xa, ya)
		cvs.setCell(xa, ya, ch, fg, curCell.Bg)
	}
}

// DrawCellBg draws a cell's background, preserving the foreground content.
func (cvs *Canvas) DrawCellBg(x, y RelCoord, bg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	xa, ya := cvs.relToAbs(x, y)
	if cvs.isIn(xa, ya) {
		curCell := cvs.cellAt(curBuf, xa, ya)
		cvs.setCell(xa, ya, curCell.Ch, curCell.Fg, bg)
	}
}

// DrawBorder draws a border in a given style, preserving the background.
func (cvs *Canvas) DrawBorder(x, y RelCoord, w, h int, borderStyle *BorderStyle, fg termbox.Attribute) {
	cvs.DrawRuneFg(x, y, borderStyle.TopLeft, fg)
	cvs.DrawRuneFg(RelCoord(int(x)+w-1), y, borderStyle.TopRight, fg)
	cvs.DrawRuneFg(x, RelCoord(int(y)+h-1), borderStyle.BottomLeft, fg)
	cvs.DrawRuneFg(RelCoord(int(x)+w-1), RelCoord(int(y)+h-1), borderStyle.BottomRight, fg)

	cvs.DrawHorizLine(x+1, y, w-2, borderStyle.Top, fg)
	cvs.DrawHorizLine(x+1, RelCoord(int(y)+h-1), w-2, borderStyle.Bottom, fg)
	cvs.DrawVertLine(x, y+1, h-2, borderStyle.Left, fg)
	cvs.DrawVertLine(RelCoord(int(x)+w-1), y+1, h-2, borderStyle.Right, fg)
}

// DrawStringFg draws a string in a given style, preserving the background.
func (cvs *Canvas) DrawStringFg(x, y RelCoord, str string, fg termbox.Attribute) {
	chars := []rune(str)
	for i, c := range chars {
		cvs.DrawRuneFg(RelCoord(int(x)+i), y, c, fg)
	}
}

func (cvs *Canvas) setCell(xa, ya AbsCoord, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(int(xa), int(ya), ch, fg, bg)
}

func (cvs *Canvas) cellAt(buf []termbox.Cell, x, y AbsCoord) termbox.Cell {
	w, _ := termbox.Size()
	lineOffset := AbsCoord(w * int(y))
	cellOffset := lineOffset + x
	return buf[cellOffset]
}
