package clikit

import "sort"

//==============================
// P O S I T I O N A L M O D E L
//==============================

// Coordinate represents the position of a component on the screen, measured
// in character cells.
type Coordinate struct {
	// The horizontal component of position.
	x RelCoord
	// The vertical component of position.
	y RelCoord
	// Used to order components from back to front in case of overlapping
	// components.
	z int
}

// X returns the coordinate's horizontal component.
func (c *Coordinate) X() RelCoord {
	return c.x
}

// SetX sets the coordinate's horizontal component.
func (c *Coordinate) SetX(x RelCoord) {
	c.x = x
}

// Y returns the coordinate's vertical component.
func (c *Coordinate) Y() RelCoord {
	return c.y
}

// SetY sets the coordinate's vertical component.
func (c *Coordinate) SetY(y RelCoord) {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Stretch adjusts a dimension's value, respecting min and max constraints
func (c *Dimension) Stretch(target int) {
	if target == c.value {
		return
	} else if target > c.value {
		c.value = min(target, c.max)
	} else if target < c.value {
		c.value = max(target, c.min)
	}
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
	Stretch(w, h int)
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

// Stretch adjusts the model's dimensions, respecting min and max constraints
func (m *DefaultPositionalModel) Stretch(w, h int) {
	m.Width().Stretch(w)
	m.Height().Stretch(h)
}

//==============================
// C O M P O N E N T
//==============================

// Component is any UI widget that can be drawn on the screen.
type Component interface {
	// Draw this Component to the screen.
	Draw(cvs Canvas)
	// PositionalModel returns the component's PositionalModel for use in layout management
	PositionalModel() PositionalModel
	// Stretch adjusts the Component's dimensions, typically respecting min and max constraints
	Stretch(w, h int)
}

//==============================
// C O N T A I N E R
//==============================

// Container is any UI widget that can contain other widgets.
type Container interface {
	Component

	// Add a Child component to this Container.
	Add(component Component, hint LayoutHint)

	// Arrange adjusts the relative positioning of child Components according to the needs
	// of a Layout Manager.
	Arrange()
}

// ChildList is list of a component's children
type ChildList []LayoutHintedComponent

// DefaultContainerModel is the base struct for all containers.
type DefaultContainerModel struct {
	DefaultPositionalModel
	ChildComponents ChildList
}

// ContainerModel is the base type for all container models.
type ContainerModel interface {
	PositionalModel
	Add(component Component, hint LayoutHint)
	Children() ChildList
}

// Add adds a component to a container.
func (cm *DefaultContainerModel) Add(component Component, hint LayoutHint) {
	cm.ChildComponents = append(cm.ChildComponents, LayoutHintedComponent{
		Component:  component,
		LayoutHint: hint,
	})
}

// Children returns the child Components of this container in the order they
// were added.
func (cm *DefaultContainerModel) Children() ChildList {
	return cm.ChildComponents
}

// InZorder non-destructively sorts a ChildList into z-order.
func (slice ChildList) InZorder() ChildList {
	newList := make(ChildList, slice.Len())
	copy(newList, slice)
	sort.Sort(newList)
	return newList
}

func (slice ChildList) Len() int {
	return len(slice)
}

func (slice ChildList) Less(i, j int) bool {
	return slice[i].Component.PositionalModel().Position().Z() < slice[j].Component.PositionalModel().Position().Z()
}

func (slice ChildList) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
