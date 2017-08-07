package clikit

import termbox "github.com/nsf/termbox-go"

// FillBgBox fills a box with a background color, preserving the foreground contents.
func FillBgBox(x, y, w, h int, bg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	for xi := x; xi < x+w; xi++ {
		for yi := y; yi < y+h; yi++ {
			curCell := cellAt(curBuf, xi, yi)
			termbox.SetCell(xi, yi, curCell.Ch, curCell.Fg, bg)
		}
	}
}

// FillFgBox fills a box with a rune, preserving the background.
func FillFgBox(x, y, w, h int, ch rune, fg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	for xi := x; xi < x+w; xi++ {
		for yi := y; yi < y+h; yi++ {
			curCell := cellAt(curBuf, xi, yi)
			termbox.SetCell(xi, yi, ch, fg, curCell.Bg)
		}
	}
}

// DrawHorizLine draws a horizontal line with a given rune in the foreground,
// preserving the background.
func DrawHorizLine(x, y, len int, ch rune, fg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	for xi := x; xi < x+len; xi++ {
		curCell := cellAt(curBuf, xi, y)
		termbox.SetCell(xi, y, ch, fg, curCell.Bg)
	}
}

// DrawCappedHorizLine draws a horizontal line with a given rune and endcaps in the foreground,
// preserving the background.
func DrawCappedHorizLine(x, y, len int, left, middle, right rune, fg termbox.Attribute) {
	DrawRuneFg(x, y, left, fg)
	DrawHorizLine(x+1, y, len-2, middle, fg)
	DrawRuneFg(x+len-1, y, right, fg)
}

// DrawHorizLineStyle draws a horizontal line with a given style in the foreground,
// preserving the background.
func DrawHorizLineStyle(x, y, len int, lineStyle *LineStyle, fg termbox.Attribute) {
	DrawCappedHorizLine(x, y, len, lineStyle.LeftCap, lineStyle.Horiz, lineStyle.RightCap, fg)
}

// DrawVertLine draws a horizontal line with a given rune in the foreground,
// preserving the background.
func DrawVertLine(x, y, len int, ch rune, fg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	for yi := y; yi < y+len; yi++ {
		curCell := cellAt(curBuf, x, yi)
		termbox.SetCell(x, yi, ch, fg, curCell.Bg)
	}
}

// DrawCappedVertLine draws a vertical line with a given rune and endcaps in the foreground,
// preserving the background.
func DrawCappedVertLine(x, y, len int, top, middle, bottom rune, fg termbox.Attribute) {
	DrawRuneFg(x, y, top, fg)
	DrawVertLine(x, y+1, len-2, middle, fg)
	DrawRuneFg(x, y+len-1, bottom, fg)
}

// DrawVertLineStyle draws a vertical line with a given style in the foreground,
// preserving the background.
func DrawVertLineStyle(x, y, len int, lineStyle *LineStyle, fg termbox.Attribute) {
	DrawCappedVertLine(x, y, len, lineStyle.TopCap, lineStyle.Vert, lineStyle.BottomCap, fg)
}

// DrawRuneFg draws a rune in the foreground, preserving the background.
func DrawRuneFg(x, y int, ch rune, fg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	curCell := cellAt(curBuf, x, y)
	termbox.SetCell(x, y, ch, fg, curCell.Bg)
}

// DrawCellBg draws a cell's background, preserving the foreground content.
func DrawCellBg(x, y int, bg termbox.Attribute) {
	curBuf := termbox.CellBuffer()
	curCell := cellAt(curBuf, x, y)
	termbox.SetCell(x, y, curCell.Ch, curCell.Fg, bg)
}

// DrawBorder draws a border in a given style, preserving the background.
func DrawBorder(x, y, w, h int, borderStyle *BorderStyle, fg termbox.Attribute) {
	DrawRuneFg(x, y, borderStyle.TopLeft, fg)
	DrawRuneFg(x+w-1, y, borderStyle.TopRight, fg)
	DrawRuneFg(x, y+h-1, borderStyle.BottomLeft, fg)
	DrawRuneFg(x+w-1, y+h-1, borderStyle.BottomRight, fg)

	DrawHorizLine(x+1, y, w-2, borderStyle.Top, fg)
	DrawHorizLine(x+1, y+h-1, w-2, borderStyle.Bottom, fg)
	DrawVertLine(x, y+1, h-2, borderStyle.Left, fg)
	DrawVertLine(x+w-1, y+1, h-2, borderStyle.Right, fg)
}

// DrawStringFg draws a string in a given style, preserving the background.
func DrawStringFg(x, y int, str string, fg termbox.Attribute) {
	chars := []rune(str)
	for i, c := range chars {
		DrawRuneFg(x+i, y, c, fg)
	}
}

func cellAt(buf []termbox.Cell, x, y int) termbox.Cell {
	w, _ := termbox.Size()
	lineOffset := w * y
	cellOffset := lineOffset + x
	return buf[cellOffset]
}
