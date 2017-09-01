package main

import (
	"github.com/nsf/termbox-go"
	"github.com/rmichela/clikit"
)

func main() {
	chkerr(clikit.Init(termbox.OutputNormal))

	label1 := clikit.NewLabel("Hello World")
	label1.Model.Position().SetX(3)
	label1.Model.Position().SetY(1)
	label1.Model.Position().SetZ(1)
	label1.Model.Width().SetValue(18)
	label1.Model.Height().SetValue(4)
	label1.StyleLabel = clikit.StyleBlack
	label1.StyleBackground = clikit.StyleRed

	label2 := clikit.NewLabel("Goodbye World")
	label2.Model.Position().SetX(2)
	label2.Model.Position().SetY(3)
	label2.Model.Position().SetZ(2)
	label2.Model.Width().SetValue(13)
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
	panel.Add(label2, clikit.LayoutHintNone)
	panel.Add(label1, clikit.LayoutHintNone)

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
