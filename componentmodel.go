package clikit

//==============================
// P O S I T I O N A L M O D E L
//==============================

// Coordinate represents the position of a component on the screen, measured
// in character cells.
type Coordinate struct {
	// The horizontal component of position.
	x int
	// The vertical component of position.
	y int
	// Used to order components from back to front in case of overlapping
	// components.
	z int
}

// X returns the coordinate's horizontal component.
func (c *Coordinate) X() int {
	return c.x
}

// SetX sets the coordinate's horizontal component.
func (c *Coordinate) SetX(x int) {
	c.x = x
}

// Y returns the coordinate's vertical component.
func (c *Coordinate) Y() int {
	return c.y
}

// SetY sets the coordinate's vertical component.
func (c *Coordinate) SetY(y int) {
	c.y = y
}

// Z returns the coordinate's ordering from front to back.
func (c *Coordinate) Z() int {
	return c.z
}

// SetZ sets the coordinate's ordering from front to back.
func (c *Coordinate) SetZ(z int) {
	c.z = z
}

// Dimension represents a bounded measurement of a component, like width and
// height. Layout managers use dimensions when resizing components at runtime.
type Dimension struct {
	// The dimension's magnitude. Layout managers may modify this value at
	// runtime.
	value int
	// The dimension's minimum value, with respect to layout management.
	min int
	// The dimension's maximum value, with respect to layout management.
	max int
}

// Value returns the dimension's magnitude. Layout managers may modify this value at
// runtime.
func (c *Dimension) Value() int {
	return c.value
}

// SetValue sets the dimension's magnitude. Layout managers may modify this value at
// runtime.
func (c *Dimension) SetValue(value int) {
	c.value = value
}

// Min returns the dimension's minimum value, with respect to layout management.
func (c *Dimension) Min() int {
	return c.min
}

// SetMin sets the dimension's minimum value, with respect to layout management.
func (c *Dimension) SetMin(min int) {
	c.min = min
}

// Max returns the dimension's maximum value, with respect to layout management.
func (c *Dimension) Max() int {
	return c.max
}

// SetMax sets the dimension's maximum value, with respect to layout management.
func (c *Dimension) SetMax(max int) {
	c.max = max
}

// DefaultPositionalModel is the base struct for all component models.
type DefaultPositionalModel struct {
	position Coordinate
	width    Dimension
	height   Dimension
}

// PositionalModel is the base interface for all component models
type PositionalModel interface {
	Position() *Coordinate
	Width() *Dimension
	Height() *Dimension
}

// Position returns the model's screen coordinates.
func (m *DefaultPositionalModel) Position() *Coordinate {
	return &m.position
}

// Width returns the model's width dimension.
func (m *DefaultPositionalModel) Width() *Dimension {
	return &m.width
}

// Height returns the model's height dimension.
func (m *DefaultPositionalModel) Height() *Dimension {
	return &m.height
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

	// Arrange adjusts the relative positioning of child Components according to the needs
	// of a Layout Manager.
	Arrange()
}

// DefaultContainerModel is the base struct for all containers.
type DefaultContainerModel struct {
	DefaultPositionalModel
	children []layoutComponent
}

// ContainerModel is the base type for all container models.
type ContainerModel interface {
	PositionalModel
	Add(component Component, hint LayoutHint)
}

// Add adds a component to a container.
func (cm *DefaultContainerModel) Add(component Component, hint LayoutHint) {
	cm.children = append(cm.children, layoutComponent{
		component: component,
		hint:      hint,
	})
}
