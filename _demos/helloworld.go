package main

import (
	"github.com/nsf/termbox-go"
	"github.com/rmichela/clikit"
)

func main() {
	chkerr(clikit.Init(termbox.OutputNormal))

	label1 := clikit.NewLabel("The WORLD!")
	label1.Model.Position().SetX(3)
	label1.Model.Position().SetY(1)
	label1.Model.Position().SetZ(1)
	label1.Model.Width().SetValue(18)
	label1.Model.Height().SetValue(4)
	label1.StyleLabel = clikit.StyleBlack
	label1.StyleBackground = clikit.StyleRed

	label2 := clikit.NewLabel("Hello World")
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

	label3 := clikit.NewLabel("AAA")
	label3.StyleBackground = clikit.StyleRed

	label4 := clikit.NewLabel("AAA")
	label4.StyleBackground = clikit.StyleGreen

	panel2 := clikit.NewPanel()
	panel2.Model.Position().SetX(40)
	panel2.Model.Position().SetY(40)
	panel2.Model.Width().SetValue(50)
	panel2.Model.Height().SetValue(8)
	panel2.Model.SetBorder(&clikit.BorderStyleDouble)
	panel2.Layout = clikit.NewRowLayout(1)
	panel2.Add(label3, clikit.LayoutHintNone)
	panel2.Add(label4, clikit.LayoutHintNone)

	screen := clikit.NewScreen()
	screen.Add(panel, clikit.LayoutHintNone)
	screen.Add(panel2, clikit.LayoutHintNone)
	screen.Refresh()

	termbox.PollEvent()

	label2.Model.SetText("Goodbye World")
	screen.Refresh()

	termbox.PollEvent()

	clikit.Close()
}

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
