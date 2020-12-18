package tile

// An object that can be rendered into a display simpleTile
type Renderable interface {
	// Render to simpleTile
	Render(rules ...RenderRule) Tile
}

// Constructor
func NewRenderable(content interface{}) Renderable {
	switch val := content.(type) {
	case Renderable:
		return val
	default:
		return NewTile(content)
	}
}
