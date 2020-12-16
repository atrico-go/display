package xy

import "fmt"

type Offset interface {
	X() int
	Y() int
	Add(rhs Offset) Offset
}

func NewOffset(x, y int) Offset {
	return offset{x: x, y: y}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type offset struct {
	x, y int
}

func (o offset) String() string {
	return fmt.Sprintf("<%d,%d>", o.x, o.y)
}

func (o offset) X() int {
	return o.x
}

func (o offset) Y() int {
	return o.y
}

func (o offset) Add(rhs Offset) Offset {
	return offset{o.x + rhs.X(), o.y + o.Y()}
}
