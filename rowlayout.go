package clikit

//=================================
// F L O W   L A Y O U T
//=================================

// NewRowLayout constructs a new row layout manager.
// A row layout places components horizontally in a line inside a container.
// Row layout will try to stretch components within their constraints to try
// to fill the row.
func NewRowLayout(hPad int) *LayoutManager {
	lm := new(LayoutManager)
	lm.Manager = rowLayoutFunc
	lm.Consraint.HPad = hPad
	return lm
}

func rowLayoutFunc(components []LayoutHintedComponent, constraint LayoutManagerConstraint) {
	// idealWidth := len(components) / constraint.Width

}
