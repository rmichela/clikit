package clikit

// DefaultPanelModel is the base struct for Panel component models.
type DefaultPanelModel struct {
	DefaultContainerModel

	border *BorderStyle
	title  string
}

// PanelModel is the model interface for Panel component models.
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
	Model           PanelModel
	Layout          *LayoutManager
	StyleBackground string
	StyleBorder     string
	StyleLabel      string
}

// NewPanel constructs a new Panel component.
func NewPanel() *Panel {
	p := new(Panel)
	p.Model = new(DefaultPanelModel)
	p.Layout = NewFixedLayout()
	p.StyleBackground = StyleComponentBackground
	p.StyleBorder = StyleComponentBorder
	p.StyleLabel = StyleComponentLabel
	return p
}

// Add adds a child component to the Panel.
func (p *Panel) Add(component Component, hint LayoutHint) {
	p.Model.Add(component, hint)
}

// Arrange adjusts the relative positioning of child Components according to the needs
// of a Layout Manager.
func (p *Panel) Arrange() {
	if p.Model.Border() != nil {
		// constrain the layout two cells smaller in each dimension to accomidate the border
		p.Layout.Resize(p.Model.Width().Value()-2, p.Model.Height().Value()-2)
		p.Layout.Apply(p.Model.Children())
		// nudge the children by one cell in each dimension to accomidate for the border
		p.Layout.Nudge(p.Model.Children(), 1, 1)
	} else {
		p.Layout.Resize(p.Model.Width().Value(), p.Model.Height().Value())
		p.Layout.Apply(p.Model.Children())
	}

	// also arrange child containers
	for _, child := range p.Model.Children() {
		c, ok := child.Component.(Container)
		if ok {
			c.Arrange()
		}
	}
}

// PositionalModel returns the component's PositionalModel for use in layout management
func (p *Panel) PositionalModel() PositionalModel {
	return p.Model
}

// Draw this Panel on the screen.
func (p *Panel) Draw(cvs Canvas) {
	// Draw background and border
	cvs.FillBgBox(
		0, 0,
		p.Model.Width().Value(),
		p.Model.Height().Value(),
		currentPalate.ResolveBg(p.StyleBackground))
	if p.Model.Border() != nil {
		cvs.DrawBorder(
			0, 0,
			p.Model.Width().Value(),
			p.Model.Height().Value(),
			p.Model.Border(),
			currentPalate.ResolveFg(p.StyleBorder))
	}
	// Draw title
	if p.Model.Title() != "" {
		cvs.DrawRuneFg(
			2, 0,
			'[',
			currentPalate.ResolveFg(p.StyleBorder))

		title := " " + p.Model.Title() + " "
		cvs.DrawStringFg(
			3, 0,
			title,
			currentPalate.ResolveFg(p.StyleLabel))
		cvs.DrawRuneFg(
			RelCoord(len(title)+3), 0,
			']',
			currentPalate.ResolveFg(p.StyleBorder))
	}

	// Draw children
	for _, child := range p.Model.Children().InZorder() {
		child.Component.Draw(cvs.ForChild(child.Component.PositionalModel()))
	}
}
