package tile

import "github.com/atrico-go/display/cells"

type RenderRule interface {
	// Process cells with this rule
	Process(cells cells.Cells) cells.Cells
}
