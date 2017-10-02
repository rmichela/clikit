package clikit

// Spacer is an invisible component used to take up space in layouts.
type Spacer struct {
	Model PositionalModel
}

// NewSpacer constructs a new spacer component.
func NewSpacer(w, h int) *Spacer {
	s := new(Spacer)
	s.Model = new(DefaultPositionalModel)
	s.Model.Width().SetValue(w)
	s.Model.Height().SetValue(h)
	return s
}

// Draw a label component to the screen.
func (s *Spacer) Draw(cvs Canvas) {
	// Don't draw anything. Spacer's
}

// PositionalModel returns the component's PositionalModel for use in layout management
func (s *Spacer) PositionalModel() PositionalModel {
	return s.Model
}

// Stretch adjusts the component's dimensions, typically respecting min and max constraints
func (s *Spacer) Stretch(w, h int) {
	s.PositionalModel().Stretch(w, h)
}
