package clikit

//==================================
// L A Y O U T   M A N A G E M E N T
//==================================

// LayoutHint defines a type for providing positioning hints to a layout manager
type LayoutHint string

// Constant hints for layout managers.
const (
	LayoutHintNone   LayoutHint = "none"
	LayoutHintLeft   LayoutHint = "left"
	LayoutHintRignt  LayoutHint = "right"
	LayoutHintTop    LayoutHint = "top"
	LayoutHintBottom LayoutHint = "bottom"
	LayoutHintCenter LayoutHint = "center"
)

// LayoutHintedComponent binds a component to a layout hint
type LayoutHintedComponent struct {
	Component  Component
	LayoutHint LayoutHint
}

// LayoutManagerFunc defines a layout manager function.
type LayoutManagerFunc func(components []LayoutHintedComponent, constraint LayoutManagerConstraint)

// LayoutManagerConstraint defines the constraints enforced by a LayoutManagerFunc.
type LayoutManagerConstraint struct {
	Width  int
	Height int
}

// LayoutManager bundles together a LayoutManagerFunc with its associated LayoutManagerConstraint.
type LayoutManager struct {
	Consraint LayoutManagerConstraint
	Manager   LayoutManagerFunc
}

// Apply executes the layout manager over a collection of components, respecting a set of constraints.
func (lm *LayoutManager) Apply(components []LayoutHintedComponent) {
	lm.Manager(components, lm.Consraint)
}

// Nudge adjusts the position of a set of components. Used for offsetting Components within
// Container after the layout has been applied.
func (lm *LayoutManager) Nudge(components []LayoutHintedComponent, x, y RelCoord) {
	for _, lhc := range components {
		p := lhc.Component.PositionalModel().Position()
		p.SetX(p.X() + x)
		p.SetY(p.Y() + y)
	}
}

// Resize resizes the layout manager's height and width constraints.
func (lm *LayoutManager) Resize(w, h int) {
	lm.Consraint.Width = w
	lm.Consraint.Height = h
}

//=================================
// F I X E D   L A Y O U T
//=================================

// NewFixedLayout constructs a new fixed layout manager. A fixed layout places components wherever
// the user set them to be.
func NewFixedLayout() *LayoutManager {
	lm := new(LayoutManager)
	lm.Manager = fixedLayoutFunc
	return lm
}

func fixedLayoutFunc(components []LayoutHintedComponent, constraint LayoutManagerConstraint) {
	// Don't do anything. This layout places components wherever they were set to be.
}
