package main

import (
	"github.com/nsf/termbox-go"
	"github.com/rmichela/clikit"
)

func main() {
	chkerr(clikit.Init(termbox.OutputNormal))

	label1 := clikit.NewLabel("Hello World")
	label1.Model.Position().SetX(2)
	label1.Model.Position().SetY(2)
	label1.Model.Width().SetValue(8)
	label1.StyleLabel = clikit.StyleBlack
	label1.StyleBackground = clikit.StyleCyan

	panel := clikit.NewPanel()
	panel.Model.SetTitle("Hello World")
	panel.Model.Position().SetX(40)
	panel.Model.Position().SetY(20)
	panel.Model.Width().SetValue(30)
	panel.Model.Height().SetValue(8)
	panel.Model.SetBorder(&clikit.BorderStyleSingle)

	cvs := clikit.NewCanvasFullScreen()

	label1.Draw(&cvs)
	panel.Draw(&cvs)

	termbox.Flush()
	termbox.PollEvent()
	clikit.Close()
}

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
