package xy

import "fmt"

type Size interface {
	X() int
	Y() int
	Add(rhs Size) Size
}

func NewSize(x, y int) Size {
	return size{x: x, y: y}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type size struct {
	x, y int
}

func (c size) String() string {
	return fmt.Sprintf("[%d,%d]", c.x, c.y)
}

func (c size) X() int {
	return c.x
}

func (c size) Y() int {
	return c.y
}

func (c size) Add(rhs Size) Size {
	return size{c.x + rhs.X(), c.y + rhs.Y()}
}
