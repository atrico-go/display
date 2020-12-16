package unit_tests

import (
	"testing"

	"github.com/atrico-go/display"
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

func Test_Panel_NoChildren(t *testing.T) {
	// Arrange

	// Act
	panel := display.NewPanel()

	// Assert
	assertRenderable(t, panel)
}

func Test_Panel_SingleChild(t *testing.T) {
	// Arrange
	content := tile.NewTile([]string{"One", "Two", "Three"})

	// Act
	panel := display.NewPanel(display.PositionedRenderable{Renderable: content, Position: xy.Origin})

	// Assert
	assertRenderable(t, panel, content.StringMl()...)
}

func Test_Panel_OverlayNoOffsetNoSizeIncrease(t *testing.T) {
	// Arrange
	content1 := tile.NewTile([]string{"One", "Two", "Three"})
	content2 := tile.NewTile([]string{"x", "y", "z"})

	// Act
	panel := display.NewPanel(
		display.PositionedRenderable{Renderable: content1, Position: xy.Origin},
		display.PositionedRenderable{Renderable: content2, Position: xy.Origin})

	// Assert
	//goland:noinspection ALL - test data not expected to be valid words
	assertRenderable(t, panel, "xne", "ywo", "zhree")
}

func Test_Panel_OverlayNoOffsetWithSizeIncrease(t *testing.T) {
	// Arrange
	content1 := tile.NewTile([]string{"One", "Two", "Three"})
	content2 := tile.NewTile([]string{"x", "yyyy", "z", "newline"})

	// Act
	panel := display.NewPanel(
		display.PositionedRenderable{Renderable: content1, Position: xy.Origin},
		display.PositionedRenderable{Renderable: content2, Position: xy.Origin})

	// Assert
	//goland:noinspection ALL - test data not expected to be valid words
	assertRenderable(t, panel, "xne", "yyyy", "zhree", "newline")
}

func Test_Panel_OverlayWithPositiveOffset(t *testing.T) {
	// Arrange
	content1 := tile.NewTile([]string{"One", "Two", "Three"})
	content2 := tile.NewTile([]string{"x", "y", "z"})

	// Act
	panel := display.NewPanel(
		display.PositionedRenderable{Renderable: content1, Position: xy.Origin},
		display.PositionedRenderable{Renderable: content2, Position: xy.NewPosition(1, 1)})

	// Assert
	assertRenderable(t, panel, "One", "Txo", "Tyree", " z")
}

func Test_Panel_OverlayWithNegativeOffset(t *testing.T) {
	// Arrange
	content1 := tile.NewTile([]string{"One", "Two", "Three"})
	content2 := tile.NewTile([]string{"x", "y", "z"})

	// Act
	panel := display.NewPanel(
		display.PositionedRenderable{Renderable: content1, Position: xy.Origin},
		display.PositionedRenderable{Renderable: content2, Position: xy.NewPosition(-1, -1)})

	// Assert
	assertRenderable(t, panel, "x   ", "yOne", "zTwo", " Three")
}

func Test_Panel_Intersections(t *testing.T) {
	// Arrange
	horizontal := 'h'
	vertical := 'v'
	rules := display.NewIntersectionRuleBuilder().
		AddIntersection(nil, &vertical, nil, &horizontal, 'A'). // top-left
		AddIntersection(nil, &vertical, &horizontal, nil, 'B'). // top-right
		AddIntersection(&vertical, nil, nil, &horizontal, 'C'). // bottom-left
		AddIntersection(&vertical, nil, &horizontal, nil, 'D'). // bottom-right
		AddIntersection(nil, &vertical, &horizontal, &horizontal, 'a'). // T-top
		AddIntersection(&vertical, nil, &horizontal, &horizontal, 'b'). // T-bottom
		AddIntersection(&vertical, &vertical, nil, &horizontal, 'c'). // T-left
		AddIntersection(&vertical, &vertical, &horizontal, nil, 'd'). // T-right
		Build()

	// Act
	horLine := display.NewHorizontalLine(horizontal, 3)
	vertLine := display.NewVerticalLine(vertical, 3)
	panel := display.NewPanel(
		display.PositionedRenderable{Position: xy.NewPosition(1, 0), Renderable: horLine},
		display.PositionedRenderable{Position: xy.NewPosition(1, 2), Renderable: horLine},
		display.PositionedRenderable{Position: xy.NewPosition(1, 4), Renderable: horLine},
		display.PositionedRenderable{Position: xy.NewPosition(0, 1), Renderable: vertLine},
		display.PositionedRenderable{Position: xy.NewPosition(2, 1), Renderable: vertLine},
		display.PositionedRenderable{Position: xy.NewPosition(4, 1), Renderable: vertLine})

	// Assert
	assertTile(t, panel.Render(rules),
		"AhahB",
		"v v v",
		"chvhd",
		"v v v",
		"ChbhD")
}
