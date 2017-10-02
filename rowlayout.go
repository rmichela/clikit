package clikit

//=================================
// F L O W   L A Y O U T
//=================================

// NewRowLayout constructs a new row layout manager.
// A row layout places components horizontally in a line inside a container.
// Row layout will try to stretch components horizontally within their constraints to try
// to fill the row. Components will be stretched vertically within their constraints to try
// to fill the container.
func NewRowLayout(hPad int) *LayoutManager {
	lm := new(LayoutManager)
	lm.Manager = rowLayoutFunc
	lm.Consraint.HPad = hPad
	return lm
}

func rowLayoutFunc(components []LayoutHintedComponent, constraint LayoutManagerConstraint) {
	allPaddingWidth := (len(components) - 1) * constraint.HPad
	idealWidth := len(components) / (constraint.Width - allPaddingWidth)
	idealHeight := constraint.Height

	// Resize
	for _, c := range components {
		c.Component.Stretch(idealWidth, idealHeight)
	}

	// Reposition
	offset := 0
	for _, c := range components {
		c.Component.PositionalModel().Position().SetX(RelCoord(offset))
		offset += c.Component.PositionalModel().Width().Value()
		offset += constraint.HPad
	}
}
