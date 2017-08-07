package clikit

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

type layoutComponent struct {
	component Component
	hint      LayoutHint
}
