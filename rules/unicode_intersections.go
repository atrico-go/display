package rules

import (
	"github.com/atrico-go/console/ansi"
	"github.com/atrico-go/console/box_drawing"

	"github.com/atrico-go/display/cells"
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

var UnicodeIntersections tile.RenderRule = unicodeIntersections{}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type unicodeIntersections struct{}

func (u unicodeIntersections) Process(cellz cells.Cells) cells.Cells {
	newCells := make(cells.Cells, len(cellz))
	origin, size := cells.GetOriginAndSize(cellz)
	for _, pos := range xy.AllPositions(origin, size) {
		matchCount := 0
		requiredChar := box_drawing.BoxParts{}
		checkAdjacentCell(cellz, pos, up, &requiredChar, &matchCount)
		checkAdjacentCell(cellz, pos, down, &requiredChar, &matchCount)
		checkAdjacentCell(cellz, pos, left, &requiredChar, &matchCount)
		checkAdjacentCell(cellz, pos, right, &requiredChar, &matchCount)
		if newChar, ok := box_drawing.GetBoxCharMixed(requiredChar); ok && matchCount > 1 {
			newCells[pos] = cells.NewCell(newChar, ansi.NoAttributes)
		} else if old, ok := cellz[pos]; ok {
			newCells[pos] = old
		}
	}
	return newCells
}

type dir int

const (
	up    dir = iota
	down  dir = iota
	left  dir = iota
	right dir = iota
)

func checkAdjacentCell(cellz cells.Cells, pos xy.Position, dir dir, requiredChar *box_drawing.BoxParts, matchCount *int) {
	cell, ok := cellz[getPosOffset(pos, dir)]
	if ok && cell.Flags().HasFlag(cells.Line) {
		if lineChar, ok := box_drawing.Lookup(cell.Char()); ok {
			setRequiredChar(requiredChar, matchCount, lineChar, dir)
		}
	}
}

func getPosOffset(position xy.Position, dir dir) xy.Position {
	switch dir {
	case up:
		return position.Up(1)
	case down:
		return position.Down(1)
	case left:
		return position.Left(1)
	case right:
		return position.Right(1)
	}
	panic("Invalid dir value")
}

func setRequiredChar(requiredChar *box_drawing.BoxParts, matchCount *int, parts box_drawing.BoxParts, dir dir) {
	switch dir {
	case up:
		if parts.Down != box_drawing.BoxNone {
			requiredChar.Up = parts.Down
			*matchCount++
		}
	case down:
		if parts.Up != box_drawing.BoxNone {
			requiredChar.Down = parts.Up
			*matchCount++
		}
	case left:
		if parts.Right != box_drawing.BoxNone {
			requiredChar.Left = parts.Right
			*matchCount++
		}
	case right:
		if parts.Left != box_drawing.BoxNone {
			requiredChar.Right = parts.Left
			*matchCount++
		}
	}
}
