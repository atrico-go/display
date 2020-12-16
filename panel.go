package display

import (
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

type PositionedRenderable struct {
	Position xy.Position
	tile.Renderable
}

// Constructors
func NewPanel(content ...PositionedRenderable) tile.Renderable {
	return panel{content}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type panel struct {
	children []PositionedRenderable
}

// Renderable
func (p panel) Render(rules ...tile.RenderRule) tile.Tile {
	canvas := NewCanvas()
	for _, child := range p.children {
		canvas.WriteAt(child.Position, child.Render())
	}
	return canvas.Render(rules...)
}
