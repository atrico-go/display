package cells

import (
	"strings"

	"github.com/atrico-go/console/ansi"

	"github.com/atrico-go/display/xy"
)

func ExpandCells(cells Cells, origin xy.Position, size xy.Size) []string {
	lines := make([]string, size.Y())
	for y := 0; y < size.Y(); y++ {
		line := strings.Builder{}
		currentAttributes := ansi.NoAttributes
		for x := 0; x < size.X(); x++ {
			if cell, ok := cells[xy.NewPosition(origin.X()+x, origin.Y()+y)]; ok {
				if cell.Attributes() != currentAttributes {
					delta := currentAttributes.CreateDeltaTo(cell.Attributes())
					line.WriteString(delta.ApplyTo(string(cell.Char())))
					currentAttributes = cell.Attributes()
				} else {
					line.WriteString(string(cell.Char()))
				}
			} else {
				delta := currentAttributes.ResetThis()
				line.WriteString(delta.ApplyTo(" "))
				currentAttributes = ansi.NoAttributes
			}
		}
		line.WriteString(currentAttributes.ResetThis().GetCodeString())
		lines[y] = line.String()
	}
	return lines
}

func GetOriginAndSize(cells Cells) (origin xy.Position, size xy.Size) {
	origin = xy.Origin
	max := xy.NewPosition(-1, -1)
	for k, _ := range cells {
		origin = xy.LowerPosition(origin, k)
		max = xy.UpperPosition(max, k)
	}
	size = xy.NewSize(max.X()-origin.X()+1, max.Y()-origin.Y()+1)
	return origin, size
}
