package clikit

type PanelModel interface {
	ContainerModel

	GetBorder() *BorderStyle
	SetBorder(border *BorderStyle)

	GetTitle() string
	SetTitle(title string)
}

type DefaultPanelModel struct {
	containerModel

	border *BorderStyle
	title  string
}

func (dpm *DefaultPanelModel) GetBorder() *BorderStyle {
	return dpm.border
}

func (dpm *DefaultPanelModel) SetBorder(border *BorderStyle) {
	dpm.border = border
}

func (dpm *DefaultPanelModel) GetTitle() string {
	return dpm.title
}

func (dpm *DefaultPanelModel) SetTitle(title string) {
	dpm.title = title
}

type Panel struct {
	Model PanelModel
}

func NewPanel() *Panel {
	p := new(Panel)
	p.Model = new(DefaultPanelModel)
	return p
}

func (p *Panel) Add(component Component, hint LayoutHint) {
	p.Model.Add(component, hint)
}

func (p *Panel) Arrange() {

}

func (p *Panel) Draw() {
	// Draw background and border
	if p.Model.GetBorder() != nil {
		FillBgBox(p.Model.GetX(), p.Model.GetY(), p.Model.GetWidth(), p.Model.GetHeight(), currentPalate.ResolveBg(StyleComponentBackground))
		DrawBorder(p.Model.GetX(), p.Model.GetY(), p.Model.GetWidth(), p.Model.GetHeight(), p.Model.GetBorder(), currentPalate.ResolveFg(StyleComponentBorder))
	}
	// Draw title
	if p.Model.GetTitle() != "" {
		DrawRuneFg(p.Model.GetX()+2, p.Model.GetY(), '[', currentPalate.ResolveFg(StyleComponentBorder))
		title := " " + p.Model.GetTitle() + " "
		DrawStringFg(p.Model.GetX()+3, p.Model.GetY(), title, currentPalate.ResolveFg(StyleComponentLabel))
		DrawRuneFg(p.Model.GetX()+len(title)+3, p.Model.GetY(), ']', currentPalate.ResolveFg(StyleComponentBorder))
	}
}
