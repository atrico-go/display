package tile

import (
	"github.com/atrico-go/console/ansi"

	"github.com/atrico-go/display/cells"
	"github.com/atrico-go/display/xy"
)

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

// Tile with "proper" cells
type tile struct {
	cells cells.Cells
	size  *xy.Size
}

func (c *tile) Size() xy.Size {
	c.lazyEvaluate()
	return *c.size
}

func (c *tile) GetCell(position xy.Position) (cell cells.Cell, ok bool) {
	c.lazyEvaluate()
	cell, ok = c.cells[position]
	return cell, ok
}

func (c *tile) StringMl(_ ...interface{}) []string {
	c.lazyEvaluate()
	return cells.ExpandCells(c.cells, xy.Origin, *c.size)
}

// Renderable (Tile is static so ignore rules)
func (c *tile) Render(_ ...RenderRule) Tile {
	return c
}

func (c *tile) lazyEvaluate() {
	if c.size == nil {
		origin, size := cells.GetOriginAndSize(c.cells)
		newCells := make(cells.Cells, size.X()*size.Y())
		// Pad missing cells and re-origin to 0,0
		for _, pos := range xy.AllPositions(origin, size) {
			newPos := pos.Offset(-origin.X(), -origin.Y())
			if cell, ok := c.cells[pos]; ok {
				newCells[newPos] = cell
			} else {
				newCells[newPos] = cells.NewCell(' ', ansi.NoAttributes)
			}
		}
		c.cells = newCells
		c.size = &size
	}
}
