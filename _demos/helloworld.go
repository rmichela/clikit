package main

import (
	"github.com/nsf/termbox-go"
	"github.com/rmichela/clikit"
)

func main() {
	chkerr(clikit.Init(termbox.OutputNormal))

	label1 := clikit.NewLabel("Hello World")
	label1.Model.Position().SetX(2)
	label1.Model.Position().SetY(1)
	label1.Model.Width().SetValue(8)
	label1.StyleLabel = clikit.StyleBlack
	label1.StyleBackground = clikit.StyleCyan

	label2 := clikit.NewLabel("Goodbye World")
	label2.Model.Position().SetX(2)
	label2.Model.Position().SetY(3)
	label2.Model.Width().SetValue(11)
	label2.StyleLabel = clikit.StyleRed
	label2.StyleBackground = clikit.StyleCyan

	panel := clikit.NewPanel()
	panel.Model.SetTitle("Hello World")
	panel.Model.Position().SetX(40)
	panel.Model.Position().SetY(20)
	panel.Model.Width().SetValue(30)
	panel.Model.Height().SetValue(8)
	panel.Model.SetBorder(&clikit.BorderStyleSingle)
	panel.StyleBackground = clikit.StyleMagenta
	panel.Add(label1, clikit.LayoutHintNone)
	panel.Add(label2, clikit.LayoutHintNone)

	screen := clikit.NewScreen()
	screen.Add(panel, clikit.LayoutHintNone)

	cvs := clikit.NewCanvasFullScreen()
	screen.Arrange()
	screen.Draw(cvs)

	termbox.Flush()
	termbox.PollEvent()
	clikit.Close()
}

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
