package main

import (
	"github.com/nsf/termbox-go"
	"github.com/rmichela/clikit"
)

func main() {
	chkerr(clikit.Init(termbox.OutputNormal))

	fg, bg := clikit.PalateDefault.Resolve(clikit.StyleRed)

	clikit.FillBgBox(0, 0, 22, 12, termbox.ColorGreen)
	clikit.FillBgBox(1, 1, 20, 10, bg)
	clikit.DrawBorder(1, 1, 20, 10, &clikit.BorderStyleDouble, fg)
	clikit.DrawHorizLineStyle(1, 5, 20, &clikit.LineStyleSingleWithDoubleCaps, fg)
	clikit.DrawVertLineStyle(5, 1, 10, &clikit.LineStyleSingleWithDoubleCaps, fg)

	panel := clikit.NewPanel()
	panel.Model.SetTitle("Hello World")
	panel.Model.Position().SetX(40)
	panel.Model.Position().SetY(20)
	panel.Model.Width().SetValue(30)
	panel.Model.Height().SetValue(8)
	panel.Model.SetBorder(&clikit.BorderStyleSingle)

	panel.Draw()
	termbox.Flush()

	termbox.PollEvent()

	clikit.Close()
}

func chkerr(err error) {
	if err != nil {
		panic(err)
	}
}
