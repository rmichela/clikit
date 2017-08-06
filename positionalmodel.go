package clikit

// positionalModel is the base struct for all component models.
type positionalModel struct {
	x int
	y int

	width    int
	maxWidth int
	minWidth int

	height    int
	maxHeight int
	minHeight int
}

type containerModel struct {
	positionalModel
}

// PositionalModel is the base type for all component models
type PositionalModel interface {
	GetX() int
	GetY() int

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
	GetMinWHeight() int
	SetMinHeight(minWidth int)
}

func (m *positionalModel) GetX() int {
	return m.x
}

func (m *positionalModel) GetY() int {
	return m.y
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
