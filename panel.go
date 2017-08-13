package clikit

// DefaultPanelModel is the base struct for Panel component models.
type DefaultPanelModel struct {
	DefaultContainerModel

	border *BorderStyle
	title  string
}

// PanelModel is the base interface for Panel component models.
type PanelModel interface {
	ContainerModel

	Border() *BorderStyle
	SetBorder(border *BorderStyle)

	Title() string
	SetTitle(title string)
}

// Border returns a pointer to the panel model's border style.
func (dpm *DefaultPanelModel) Border() *BorderStyle {
	return dpm.border
}

// SetBorder sets the panel model's border style.
func (dpm *DefaultPanelModel) SetBorder(border *BorderStyle) {
	dpm.border = border
}

// Title returns the panel model's title.
func (dpm *DefaultPanelModel) Title() string {
	return dpm.title
}

// SetTitle sets the panel model's title.
func (dpm *DefaultPanelModel) SetTitle(title string) {
	dpm.title = title
}

// Panel is a Container component with an optional border and title.
type Panel struct {
	Model PanelModel
}

// NewPanel constructs a new Panel component.
func NewPanel() *Panel {
	p := new(Panel)
	p.Model = new(DefaultPanelModel)
	return p
}

// Add adds a child component to the Panel.
func (p *Panel) Add(component Component, hint LayoutHint) {
	p.Model.Add(component, hint)
}

// Arrange adjusts the relative positioning of child Components according to the needs
// of a Layout Manager.
func (p *Panel) Arrange() {

}

// Draw this Panel on the screen.
func (p *Panel) Draw() {
	// Draw background and border
	if p.Model.Border() != nil {
		FillBgBox(p.Model.Position().X(), p.Model.Position().Y(), p.Model.Width().Value(), p.Model.Height().Value(), currentPalate.ResolveBg(StyleComponentBackground))
		DrawBorder(p.Model.Position().X(), p.Model.Position().Y(), p.Model.Width().Value(), p.Model.Height().Value(), p.Model.Border(), currentPalate.ResolveFg(StyleComponentBorder))
	}
	// Draw title
	if p.Model.Title() != "" {
		DrawRuneFg(p.Model.Position().X()+2, p.Model.Position().Y(), '[', currentPalate.ResolveFg(StyleComponentBorder))
		title := " " + p.Model.Title() + " "
		DrawStringFg(p.Model.Position().X()+3, p.Model.Position().Y(), title, currentPalate.ResolveFg(StyleComponentLabel))
		DrawRuneFg(p.Model.Position().X()+len(title)+3, p.Model.Position().Y(), ']', currentPalate.ResolveFg(StyleComponentBorder))
	}
}
