package display

import (
	"github.com/atrico-go/display/cells"
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

// Mutable display type
type Canvas interface {
	// Modify
	WriteAt(position xy.Position, content interface{}) Canvas
	// Cells
	GetCell(position xy.Position) (cell cells.Cell, ok bool)
	SetCell(position xy.Position, cell cells.Cell)
	// Renderable
	Render(rules ...tile.RenderRule) tile.Tile
}

func NewCanvas() Canvas {
	return &canvas{cells: make(map[xy.Position]cells.Cell)}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type canvas struct {
	cells cells.Cells
}

// Modify
func (c *canvas) WriteAt(position xy.Position, content interface{}) Canvas {
	// Create static tile
	tile := tile.NewTile(content)
	for _, pos := range xy.AllPositions(xy.Origin, tile.Size()) {
		if cell, ok := tile.GetCell(pos); ok {
			c.setCell(position.Offset(pos.X(), pos.Y()), cell)
		}
	}
	return c
}

// Cells
func (c *canvas) GetCell(position xy.Position) (cell cells.Cell, ok bool) {
	cell, ok = c.cells[position]
	return cell, ok
}

func (c *canvas) SetCell(position xy.Position, cell cells.Cell) {
	c.cells[position] = cell
}

// Renderable
func (c *canvas) Render(rules ...tile.RenderRule) tile.Tile {
	theCells := c.cells
	for _, rule := range rules {
		theCells = rule.Process(theCells)
	}
	// origin, size := cells.GetOriginAndSize(theCells)
	// return tile.NewTile(cells.ExpandCells(theCells, origin, size))
	return tile.NewTile(theCells)
}

func (c *canvas) setCell(position xy.Position, cell cells.Cell) {
	// Write new cell
	if cell.Char() != ' ' {
		// If not transparent, overwrite
		c.cells[position] = cell
	} else {
		// If transparent, only write if nothing already there (to keep size correct)
		if _, ok := c.cells[position]; !ok {
			c.cells[position] = cell
		}
	}
}
