package clikit

import (
	"github.com/nsf/termbox-go"
)

// Init initializes the clikit library, and the underlying termbox.
func Init() error {
	return termbox.Init()
}

// Close terminates the clikit library, and underlying termbox.
func Close() {
	termbox.Close()
}
