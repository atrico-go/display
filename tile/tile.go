package tile

import (
	"fmt"
	"strings"

	"github.com/atrico-go/console/ansi"

	"github.com/atrico-go/display/cells"
	"github.com/atrico-go/display/common"

	. "github.com/atrico-go/core"

	"github.com/atrico-go/display/xy"
)

// An static displayable rectangle
type Tile interface {
	// Size in characters
	Size() xy.Size
	// Get contents of a specific cell
	GetCell(position xy.Position) (cell cells.Cell, ok bool)
	// Representation of object as multiple lines
	StringMl(params ...interface{}) []string
	// Tile is intrinsically renderable (generally return this)
	Render(rules ...RenderRule) Tile
}

// Constructor
func NewTile(content interface{}) Tile {
	// Cast
	switch val := content.(type) {
	case Tile:
		return val
	case Renderable:
		return val.Render()
	}

	var cellz cells.Cells = nil
	var lines []string = nil
	switch val := content.(type) {
	case cells.Cells:
		cellz = val
	case []string:
		lines = val
	case []interface{}:
		for _, v := range val {
			lines = append(lines, fmt.Sprintf("%v", v))
		}
	case fmt.Stringer:
		lines = []string{val.String()}
	case StringerMl:
		lines = val.StringMl()
	default:
		lines = []string{fmt.Sprintf("%v", val)}
	}
	if lines != nil {
		cellz = make(cells.Cells)
		pos := xy.Origin
		for _, line := range lines {
			stringToCells(&cellz, line, pos)
			pos = pos.Down(1)
		}
	}
	return &tile{cellz, nil}
}

// Tile with minimum size
func NewSizedTile(size xy.Size, content interface{}) Tile {
	temporary := NewTile(content)
	if temporary.Size() == size {
		return temporary
	}
	lines := temporary.StringMl()
	len := len(lines)
	// Pad/clip height
	if len < size.Y() {
		lines = append(lines, make([]string, size.Y()-len)...)
	} else if len > size.Y() {
		lines = lines[:size.Y()]
	}
	// Pad/clip width
	lines = padOrClipLines(lines, size.X())
	return NewTile(lines)
}

// ----------------------------------------------------------------------------------------------------------------------------
// internal
// ----------------------------------------------------------------------------------------------------------------------------

func stringToCells(cellz *cells.Cells, str string, origin xy.Position) {
	for _, part := range ansi.ParseString(str) {
		for i, char := range []rune(part.String) {
			(*cellz)[origin.Right(i)] = cells.NewCell(char, part.Attributes)
		}
	}
}

func padOrClipLines(lines []string, width int) []string {
	newLines := make([]string, len(lines))
	for i, line := range lines {
		len := common.StringWidth(line)
		if width > len {
			newLines[i] = fmt.Sprintf("%s%s", line, strings.Repeat(" ", width-len))
		} else if len > width {
			newLines[i] = line[:width]
		} else {
			newLines[i] = line
		}
	}
	return newLines
}
