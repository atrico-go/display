package display

import (
	"github.com/atrico-go/console/ansi"

	"github.com/atrico-go/display/cells"
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

type IntersectionRuleBuilder interface {
	AddIntersection(above, below, left, right *rune, intersection rune) IntersectionRuleBuilder
	Build() tile.RenderRule
}

func NewIntersectionRuleBuilder() IntersectionRuleBuilder {
	return &rules{make(map[combination]rune)}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
type combination struct {
	above, below, left, right rune
}
type rules struct {
	rules map[combination]rune
}

func (r *rules) AddIntersection(above, below, left, right *rune, intersection rune) IntersectionRuleBuilder {
	comb := combination{0, 0, 0, 0}
	matchCount := 0
	if above != nil {
		comb.above = *above
		matchCount++
	}
	if below != nil {
		comb.below = *below
		matchCount++
	}
	if left != nil {
		comb.left = *left
		matchCount++
	}
	if right != nil {
		comb.right = *right
		matchCount++
	}
	if matchCount > 1 {
		r.rules[comb] = intersection
	}
	return r
}

func (r *rules) Build() tile.RenderRule {
	return *r
}

func (r rules) Process(cellz cells.Cells) cells.Cells {
	newCells := make(cells.Cells)
	origin, size := cells.GetOriginAndSize(cellz)
	for _, pos := range xy.AllPositions(origin, size) {
		comb := combination{0, 0, 0, 0}
		matchCount := 0
		cell, ok := cellz[pos.Up(1)]
		if ok && cell.Flags().HasFlag(cells.Line) {
			comb.above = cell.Char()
			matchCount++
		}
		cell, ok = cellz[pos.Down(1)]
		if ok && cell.Flags().HasFlag(cells.Line) {
			comb.below = cell.Char()
			matchCount++
		}
		cell, ok = cellz[pos.Left(1)]
		if ok && cell.Flags().HasFlag(cells.Line) {
			comb.left = cell.Char()
			matchCount++
		}
		cell, ok = cellz[pos.Right(1)]
		if ok && cell.Flags().HasFlag(cells.Line) {
			comb.right = cell.Char()
			matchCount++
		}
		if replace, ok := r.rules[comb]; ok && matchCount > 1 {
			newCells[pos] = cells.NewCell(replace, ansi.NoAttributes, cells.Line)
		} else {
			if old, ok := cellz[pos]; ok {
				newCells[pos] = old
			}
		}
	}
	return newCells
}
