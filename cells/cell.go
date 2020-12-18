package cells

import (
	"github.com/atrico-go/console/ansi"

	"github.com/atrico-go/display/xy"
)

// Single cell in a canvas
type Cell interface {
	Char() rune
	Attributes() ansi.Attributes
	Flags() Flags
}

type Cells map[xy.Position]Cell

type ModifyCells interface {
	// Get the cells (modifiable)
	GetCells() Cells
}

func NewCell(char rune, attributes ansi.Attributes, flags ...Flag) Cell {
	return cell{char, attributes, NewFlags(flags...)}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type cell struct {
	rune
	attributes ansi.Attributes
	flags      Flags
}

func (c cell) Attributes() ansi.Attributes {
	return c.attributes
}

func (c cell) Char() rune {
	return c.rune
}

func (c cell) Flags() Flags {
	return c.flags
}
