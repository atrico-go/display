package xy

import (
	"fmt"

	"github.com/atrico-go/display/common"
)

type Position interface {
	X() int
	Y() int
	Left(delta int) Position
	Right(delta int) Position
	Up(delta int) Position
	Down(delta int) Position
	Offset(x, y int) Position
	Add(size Size) Position
}

func NewPosition(x, y int) Position {
	return position{x: x, y: y}
}

var Origin = NewPosition(0, 0)

func UpperPosition(a, b Position) Position {
	return NewPosition(common.MaxInt(a.X(), b.X()), common.MaxInt(a.Y(), b.Y()))
}

func LowerPosition(a, b Position) Position {
	return NewPosition(common.MinInt(a.X(), b.X()), common.MinInt(a.Y(), b.Y()))
}

// All positions to range over
func AllPositions(origin Position, size Size) []Position {
	all := make([]Position, size.X()*size.Y())
	end := origin.Add(size)
	idx := 0
	for y := origin.Y(); y < end.Y(); y++ {
		for x := origin.X(); x < end.X(); x++ {
			all[idx] = NewPosition(x, y)
			idx++
		}
	}
	return all
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type position struct {
	x, y int
}

func (p position) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p position) X() int {
	return p.x
}

func (p position) Y() int {
	return p.y
}

func (p position) Left(delta int) Position {
	return NewPosition(p.X()-delta, p.Y())
}

func (p position) Right(delta int) Position {
	return NewPosition(p.X()+delta, p.Y())
}

func (p position) Up(delta int) Position {
	return NewPosition(p.X(), p.Y()-delta)
}

func (p position) Down(delta int) Position {
	return NewPosition(p.X(), p.Y()+delta)
}

func (p position) Offset(x, y int) Position {
	return position{p.x + x, p.y + y}
}

func (p position) Add(size Size) Position {
	return position{p.x + size.X(), p.y + size.Y()}
}
