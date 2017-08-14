package clikit

import termbox "github.com/nsf/termbox-go"

// Screen is a Container component representing the root component in a user
// interface. Only one screen is active at a time. Changing UI modes is
// accomplished by switching the active Screen.
type Screen struct {
	Model  ContainerModel
	Layout *LayoutManager
}

// NewScreen constructs a new Screen with the current dimensions of the terminal.
func NewScreen() *Screen {
	s := new(Screen)
	s.Model = new(DefaultContainerModel)
	s.Layout = NewFixedLayout()
	s.Resize(termbox.Size())
	return s
}

// Add adds a child component to the Panel.
func (s *Screen) Add(component Component, hint LayoutHint) {
	s.Model.Add(component, hint)
}

// Resize changes the size of a screen.
func (s *Screen) Resize(w, h int) {
	s.Model.Width().SetValue(w)
	s.Model.Width().SetMin(w)
	s.Model.Width().SetMax(w)

	s.Model.Height().SetValue(w)
	s.Model.Height().SetMin(w)
	s.Model.Height().SetMax(w)

	s.Arrange()
}

// Arrange adjusts the relative positioning of child Components according to the needs
// of a Layout Manager.
func (s *Screen) Arrange() {
	s.Layout.Resize(s.Model.Width().Value(), s.Model.Height().Value())
	s.Layout.Apply(s.Model.Children())

	// also arrange child containers
	for _, child := range s.Model.Children() {
		c, ok := child.Component.(Container)
		if ok {
			c.Arrange()
		}
	}
}

// PositionalModel returns the component's PositionalModel for use in layout management
func (s *Screen) PositionalModel() PositionalModel {
	return s.Model
}

// Draw this Panel on the screen.
func (s *Screen) Draw(cvs Canvas) {
	for _, child := range s.Model.Children() {
		child.Component.Draw(cvs.ForChild(child.Component.PositionalModel()))
	}
}
