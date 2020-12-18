package display

import (
	"github.com/atrico-go/console/ansi"

	"github.com/atrico-go/display/cells"
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

func NewHorizontalLine(char rune, len int) tile.Renderable {
	return horizontaLine{char, len}
}

func NewVerticalLine(char rune, len int) tile.Renderable {
	return verticaLine{char, len}

}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type line struct {
	rune
	length int
}
type horizontaLine line

func (l horizontaLine) Render(_ ...tile.RenderRule) tile.Tile {
	cellz := make(cells.Cells)
	for _, pos := range xy.AllPositions(xy.Origin, xy.NewSize(l.length, 1)) {
		cellz[pos] = cells.NewCell(l.rune, ansi.NoAttributes, cells.Line)
	}
	return tile.NewTile(cellz)
}

type verticaLine line

func (l verticaLine) Render(_ ...tile.RenderRule) tile.Tile {
	cellz := make(cells.Cells)
	for _, pos := range xy.AllPositions(xy.Origin, xy.NewSize(1, l.length)) {
		cellz[pos] = cells.NewCell(l.rune, ansi.NoAttributes, cells.Line)
	}
	return tile.NewTile(cellz)
}
