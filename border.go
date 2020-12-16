package display

import (
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

type BorderBuilder interface {
	WithContent(content tile.Renderable) BorderBuilder
	WithAbove(character rune) BorderBuilder
	WithBelow(character rune) BorderBuilder
	WithHorizontal(character rune) BorderBuilder
	WithLeft(character rune) BorderBuilder
	WithRight(character rune) BorderBuilder
	WithVertical(character rune) BorderBuilder
	Build() tile.Renderable
}

func NewBorder(content tile.Renderable, above *rune, below *rune, left *rune, right *rune) tile.Renderable {
	builder := NewBorderBuilder().WithContent(content)
	if above != nil {
		builder.WithAbove(*above)
	}
	if below != nil {
		builder.WithBelow(*below)
	}
	if left != nil {
		builder.WithLeft(*left)
	}
	if right != nil {
		builder.WithRight(*right)
	}
	return builder.Build()
}

func NewBorderBuilder() BorderBuilder {
	return &border{content: tile.NewTile("")}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type border struct {
	content tile.Renderable
	above   *rune
	below   *rune
	left    *rune
	right   *rune
}

func (b *border) WithContent(content tile.Renderable) BorderBuilder {
	b.content = content
	return b
}

func (b *border) WithAbove(character rune) BorderBuilder {
	b.above = &character
	return b
}

func (b *border) WithBelow(character rune) BorderBuilder {
	b.below = &character
	return b
}

func (b *border) WithHorizontal(character rune) BorderBuilder {
	return b.WithAbove(character).WithBelow(character)
}

func (b *border) WithLeft(character rune) BorderBuilder {
	b.left = &character
	return b
}

func (b *border) WithRight(character rune) BorderBuilder {
	b.right = &character
	return b
}

func (b *border) WithVertical(character rune) BorderBuilder {
	return b.WithLeft(character).WithRight(character)
}

func (b *border) Build() tile.Renderable {
	return *b
}

// Renderable
func (b border) Render(rules ...tile.RenderRule) tile.Tile {
	elements := make([]PositionedRenderable, 1, 5)
	elements[0] = PositionedRenderable{xy.Origin, b.content}
	main := b.content.Render()
	if b.above != nil {
		elements = append(elements, PositionedRenderable{xy.NewPosition(0, -1), NewHorizontalLine(*b.above, main.Size().X())})
	}
	if b.below != nil {
		elements = append(elements, PositionedRenderable{xy.NewPosition(0, main.Size().Y()), NewHorizontalLine(*b.below, main.Size().X())})
	}
	if b.left != nil {
		elements = append(elements, PositionedRenderable{xy.NewPosition(-1, 0), NewVerticalLine(*b.left, main.Size().Y())})
	}
	if b.right != nil {
		elements = append(elements, PositionedRenderable{xy.NewPosition(main.Size().X(), 0), NewVerticalLine(*b.right, main.Size().Y())})
	}
	return NewPanel(elements...).Render(rules...)
}
