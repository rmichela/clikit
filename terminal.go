package clikit

import (
	"errors"

	"github.com/nsf/termbox-go"
)

var currentOutputMode termbox.OutputMode
var currentPalate Palate

// Init initializes the clikit library, and the underlying termbox.
func Init(mode termbox.OutputMode) error {
	er := termbox.Init()
	if er == nil {
		if mode != termbox.OutputNormal && mode != termbox.Output256 {
			er = errors.New("Only modes OutputNormal and Output256 are supported")
		} else {
			currentOutputMode = mode
			termbox.SetOutputMode(mode)
		}

		currentPalate = PalateDefault
	}
	return er
}

// SetPalate sets the active color palate for rendering components
func SetPalate(palate Palate) {
	currentPalate = palate
}

// Close terminates the clikit library, and underlying termbox.
func Close() {
	termbox.Close()
}
