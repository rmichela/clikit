package clikit

// Screen is a Container component representing the root component in a user
// interface. Only one screen is active at a time. Changing UI modes is
// accomplished by switching the active Screen.
type Screen struct {
	Model DefaultContainerModel
}

// Add adds a child component to the Panel.
func (s *Screen) Add(component Component, hint LayoutHint) {
	s.Model.Add(component, hint)
}

// Arrange adjusts the relative positioning of child Components according to the needs
// of a Layout Manager.
func (s *Screen) Arrange() {

}

// Draw this Panel on the screen.
func (s *Screen) Draw() {

}
