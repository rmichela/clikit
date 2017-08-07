package clikit

//==============================
// P O S I T I O N A L M O D E L
//==============================

// positionalModel is the base struct for all component models.
type positionalModel struct {
	x int
	y int
	z int

	width    int
	maxWidth int
	minWidth int

	height    int
	maxHeight int
	minHeight int
}

// PositionalModel is the base type for all component models
type PositionalModel interface {
	GetX() int
	GetY() int
	GetZ() int

	SetX(x int)
	SetY(y int)
	SetZ(z int)

	GetWidth() int
	SetWidth(width int)
	GetMaxWidth() int
	SetMaxWidth(maxWidth int)
	GetMinWidth() int
	SetMinWidth(minWidth int)

	GetHeight() int
	SetHeight(width int)
	GetMaxHeight() int
	SetMaxHeight(maxWidth int)
	GetMinHeight() int
	SetMinHeight(minWidth int)
}

func (m *positionalModel) GetX() int {
	return m.x
}

func (m *positionalModel) GetY() int {
	return m.y
}

func (m *positionalModel) GetZ() int {
	return m.z
}

func (m *positionalModel) SetX(x int) {
	m.x = x
}

func (m *positionalModel) SetY(y int) {
	m.y = y
}

func (m *positionalModel) SetZ(z int) {
	m.z = z
}

func (m *positionalModel) GetWidth() int {
	return m.width
}

func (m *positionalModel) SetWidth(width int) {
	m.width = width
}

func (m *positionalModel) GetMaxWidth() int {
	return m.maxWidth
}

func (m *positionalModel) SetMaxWidth(width int) {
	m.maxWidth = width
}

func (m *positionalModel) GetMinWidth() int {
	return m.minWidth
}

func (m *positionalModel) SetMinWidth(width int) {
	m.minWidth = width
}

func (m *positionalModel) GetHeight() int {
	return m.height
}

func (m *positionalModel) SetHeight(height int) {
	m.height = height
}

func (m *positionalModel) GetMaxHeight() int {
	return m.maxHeight
}

func (m *positionalModel) SetMaxHeight(height int) {
	m.maxHeight = height
}

func (m *positionalModel) GetMinHeight() int {
	return m.minHeight
}

func (m *positionalModel) SetMinHeight(height int) {
	m.minHeight = height
}

//==============================
// C O M P O N E N T
//==============================

// Component is any UI widget that can be drawn on the screen.
type Component interface {
	// Draw this Component to the screen.
	Draw()
}

//==============================
// C O N T A I N E R
//==============================

// Container is any UI widget that can contain other widgets.
type Container interface {
	Component

	// Add a Child component to this Container, with
	Add(component Component, hint LayoutHint)

	// Adjust the relative positioning of child Components.
	Arrange()
}

type containerModel struct {
	positionalModel
	children []layoutComponent
}

// ContainerModel is the base type for all container models
type ContainerModel interface {
	PositionalModel
	Add(component Component, hint LayoutHint)
}

func (cm *containerModel) Add(component Component, hint LayoutHint) {
	cm.children = append(cm.children, layoutComponent{
		component: component,
		hint:      hint,
	})
}
