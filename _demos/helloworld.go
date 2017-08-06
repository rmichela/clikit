package main

import (
	"github.com/nsf/termbox-go"
	"github.com/rmichela/clikit"
)

func main() {
	chkerr(clikit.Init())

	clikit.FillBgBox(1, 1, 20, 10, termbox.ColorRed)
	clikit.DrawBorder(1, 1, 20, 10, clikit.BorderStyleDouble, termbox.ColorWhite)
	clikit.DrawCappedHorizLine(1, 5, 20, '╟', '─', '╢', termbox.ColorWhite)
	clikit.DrawCappedVertLine(5, 1, 10, '╤', '│', '╧', termbox.ColorWhite)
	termbox.Flush()

	termbox.PollEvent()

	clikit.Close()
}

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
