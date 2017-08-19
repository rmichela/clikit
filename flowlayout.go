package clikit

//=================================
// F L O W   L A Y O U T
//=================================

// NewFlowLayout constructs a new flow layout manager.
func NewFlowLayout(orientation LayoutOrientation) *LayoutManager {
	lm := new(LayoutManager)
	lm.Manager = flowLayoutFunc
	lm.Consraint.Orientation = orientation
	return lm
}

func flowLayoutFunc(components []LayoutHintedComponent, constraint LayoutManagerConstraint) {

}
