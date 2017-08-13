package clikit

import "github.com/mattn/go-runewidth"

// DefaultLabelModel is the base struct for Label components.
type DefaultLabelModel struct {
	DefaultPositionalModel
	text string
}

// LabelModel is the model interface for Label components
type LabelModel interface {
	PositionalModel
	Text() string
	SetText(text string)
}

// Text returns the label model's label text.
func (dlm *DefaultLabelModel) Text() string {
	return dlm.text
}

// SetText sets the label model's label text.
func (dlm *DefaultLabelModel) SetText(text string) {
	dlm.text = text
}

// Label is a textual label component.
type Label struct {
	Model           LabelModel
	StyleBackground string
	StyleLabel      string
}

// NewLabel constructs a new label component.
func NewLabel(text string) *Label {
	l := new(Label)
	l.Model = new(DefaultLabelModel)
	l.Model.SetText(text)
	l.Model.Width().SetValue(runewidth.StringWidth(text))
	l.Model.Height().SetValue(1)
	l.StyleBackground = StyleComponentBackground
	l.StyleLabel = StyleComponentLabel
	return l
}

// Draw a label component to the screen.
func (l *Label) Draw() {
	s := runewidth.Truncate(l.Model.Text(), l.Model.Width().Value(), "…")
	FillBgBox(l.Model.Position().X(), l.Model.Position().Y(), l.Model.Width().Value(), l.Model.Height().Value(), currentPalate.ResolveBg(l.StyleBackground))
	DrawStringFg(l.Model.Position().X(), l.Model.Position().Y(), s, currentPalate.ResolveFg(l.StyleLabel))
}