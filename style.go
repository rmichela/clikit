package clikit

import termbox "github.com/nsf/termbox-go"

// Standard style name constants
const (
	// Color styles
	StyleBlack   = "black"
	StyleRed     = "red"
	StyleGreen   = "green"
	StyleYellow  = "yellow"
	StyleBlue    = "blue"
	StyleMagenta = "magenta"
	StyleCyan    = "cyan"
	StyleWhite   = "white"
	// Component styles
	StyleComponentBackground = "component_background"
	StyleComponentBorder     = "component_border"
	StyleComponentLabel      = "component_label"
)

// StyleMap is a mapping of style names to termbox Attributes
type StyleMap struct {
	// A map from style name to forgreound display attribute
	StyleFG map[string]termbox.Attribute
	// A map from style name to background display attribute
	StyleBG map[string]termbox.Attribute
	// Foreground display attribute to use for unknown styles
	DefaultFG termbox.Attribute
	// Background display attribute to use for unknown styles
	DefaultBG termbox.Attribute
}

// Palate represents an overall color palate. It is made up of a pair of
// StyleMaps, one for low-fi 16-color mode, and one for hi-fi 256-color mode.
type Palate struct {
	// The stylemap to use in 16-color mode
	Style16 StyleMap
	// The stylemap to use in 256-color mode
	Style256 StyleMap
}

// PalateBlackAndWhite is a simple color palate for drawing the UI in
// only black and white.
var PalateBlackAndWhite = Palate{
	Style16: StyleMap{
		DefaultFG: termbox.ColorWhite,
		DefaultBG: termbox.ColorBlack},
	Style256: StyleMap{
		DefaultFG: termbox.ColorWhite,
		DefaultBG: termbox.ColorBlack}}

// PalateDefault implements a basic 16-color palate for both 16 and 256
// color modes.
var PalateDefault = Palate{
	Style16: StyleMap{
		DefaultFG: termbox.ColorWhite,
		DefaultBG: termbox.ColorBlack,
		StyleFG: map[string]termbox.Attribute{
			// Color Styles
			StyleBlack:   termbox.ColorBlack,
			StyleRed:     termbox.ColorRed,
			StyleGreen:   termbox.ColorGreen,
			StyleYellow:  termbox.ColorYellow,
			StyleBlue:    termbox.ColorBlue,
			StyleMagenta: termbox.ColorMagenta,
			StyleCyan:    termbox.ColorCyan,
			StyleWhite:   termbox.ColorWhite,
			// Component Styles
			StyleComponentBackground: termbox.ColorBlack,
			StyleComponentBorder:     termbox.ColorBlue,
			StyleComponentLabel:      termbox.ColorYellow,
		},
	},
}

// Resolve foreground and background attributes for a named style, returning defaults if not found.
func (p *Palate) Resolve(style string) (foreground, background termbox.Attribute) {
	var firstMap *StyleMap
	var secondMap *StyleMap

	if currentOutputMode == termbox.OutputNormal {
		firstMap = &p.Style16
		secondMap = &p.Style16
	} else if currentOutputMode == termbox.Output256 {
		firstMap = &p.Style256
		secondMap = &p.Style16
	}

	// default styles, if nothing is found
	var fg = termbox.ColorWhite
	var bg = termbox.ColorBlack

	// find the styles in the map
	var found bool
	fg, found = firstMap.StyleFG[style]
	if !found {
		fg, found = secondMap.StyleFG[style]
		if !found {
			fg = firstMap.DefaultFG
		}
	}
	bg, found = firstMap.StyleBG[style]
	if !found {
		bg, found = secondMap.StyleBG[style]
		if !found {
			bg = firstMap.DefaultBG
		}
	}

	return fg, bg
}

// ResolveFg resolves the foreground attribute for a named style, returning default if not found.
func (p *Palate) ResolveFg(style string) termbox.Attribute {
	fg, _ := p.Resolve(style)
	return fg
}

// ResolveBg resolves the background attribute for a named style, returning default if not found.
func (p *Palate) ResolveBg(style string) termbox.Attribute {
	_, bg := p.Resolve(style)
	return bg
}

// BorderStyle describes the characters that make up a drawn border.
type BorderStyle struct {
	Top         rune
	Bottom      rune
	Left        rune
	Right       rune
	TopLeft     rune
	TopRight    rune
	BottomLeft  rune
	BottomRight rune
}

// BorderStyleSingle is a border of single-stroke line drawing characters.
var BorderStyleSingle = BorderStyle{
	Top:         '─',
	Bottom:      '─',
	Left:        '│',
	Right:       '│',
	TopLeft:     '┌',
	TopRight:    '┐',
	BottomLeft:  '└',
	BottomRight: '┘',
}

// BorderStyleDouble is a border of double-stroke line drawing characters.
var BorderStyleDouble = BorderStyle{
	Top:         '═',
	Bottom:      '═',
	Left:        '║',
	Right:       '║',
	TopLeft:     '╔',
	TopRight:    '╗',
	BottomLeft:  '╚',
	BottomRight: '╝',
}

// LineStyle describes the characters that make up a drawn line with endcaps.
type LineStyle struct {
	TopCap    rune
	Vert      rune
	BottomCap rune
	LeftCap   rune
	Horiz     rune
	RightCap  rune
}

// LineStyleSingle is a line of single-stroke line drawing characters.
var LineStyleSingle = LineStyle{
	TopCap:    '│',
	Vert:      '│',
	BottomCap: '│',
	LeftCap:   '─',
	Horiz:     '─',
	RightCap:  '─',
}

// LineStyleSingleWithDoubleCaps is a line of single-stroke line drawing characters with duble-strike caps.
var LineStyleSingleWithDoubleCaps = LineStyle{
	TopCap:    '╤',
	Vert:      '│',
	BottomCap: '╧',
	LeftCap:   '╟',
	Horiz:     '─',
	RightCap:  '╢',
}

// LineStyleDouble is a line of double-stroke line drawing characters.
var LineStyleDouble = LineStyle{
	TopCap:    '║',
	Vert:      '║',
	BottomCap: '║',
	LeftCap:   '═',
	Horiz:     '═',
	RightCap:  '═',
}

// LineStyleDoubleWithSingleCaps is a line of double-stroke line drawing characters with single-strike caps.
var LineStyleDoubleWithSingleCaps = LineStyle{
	TopCap:    '╥',
	Vert:      '║',
	BottomCap: '╨',
	LeftCap:   '╞',
	Horiz:     '═',
	RightCap:  '╡',
}
