package unit_tests

import (
	"fmt"
	"testing"

	"github.com/atrico-go/console/ansi"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/display/cells"
	"github.com/atrico-go/display/tile"

	. "github.com/atrico-go/testing/assert"

	"github.com/atrico-go/display"
	"github.com/atrico-go/display/xy"
)

func Test_Canvas_Empty(t *testing.T) {
	// Arrange

	// Act
	canvas := display.NewCanvas()

	// Assert
	assertRenderable(t, canvas)
}

func Test_Canvas_WriteAtOrigin(t *testing.T) {
	// Arrange
	text := randomValues.String()
	canvas := display.NewCanvas()

	// Act
	canvas.WriteAt(xy.NewPosition(0, 0), text)

	// Assert
	assertRenderable(t, canvas, text)
}

func Test_Canvas_WriteAtOffset(t *testing.T) {
	// Arrange
	text := randomValues.String()
	canvas := display.NewCanvas()

	// Act
	canvas.WriteAt(xy.NewPosition(2, 3), text)

	// Assert
	assertRenderable(t, canvas, "", "", "", fmt.Sprintf("  %s", text))
}

func Test_Canvas_Overwrite(t *testing.T) {
	// Arrange
	text1 := []string{randomValues.String(), randomValues.String()}
	text2 := randomValues.String()
	text3 := randomValues.String()
	canvas := display.NewCanvas()
	fmt.Printf("%v,%v,%v,%v\n", text1[0], text1[1], text2, text3)

	// Act
	canvas.WriteAt(xy.NewPosition(0, 0), text1)
	canvas.WriteAt(xy.NewPosition(1, 0), text2)
	canvas.WriteAt(xy.NewPosition(2, 1), text3)

	// Assert
	exp1 := fmt.Sprintf("%s%s", text1[0][0:1], text2)
	exp2 := fmt.Sprintf("%s%s", text1[1][:2], text3)
	assertRenderable(t, canvas, exp1, exp2)
}

func Test_Canvas_EmptySizedTile(t *testing.T) {
	// Arrange
	text := tile.NewSizedTile(xy.NewSize(3, 3), "")
	canvas := display.NewCanvas()

	// Act
	canvas.WriteAt(xy.Origin, text)

	// Assert
	assertRenderable(t, canvas, "   ", "   ", "   ")
}

func Test_Canvas_GetCellEmpty(t *testing.T) {
	// Arrange
	canvas := display.NewCanvas()

	// Act
	_, ok := canvas.GetCell(xy.Origin)

	// Assert
	Assert(t).That(ok, is.False, "Nothing at this position")
}

func Test_Canvas_GetCellCharacter(t *testing.T) {
	// Arrange
	canvas := display.NewCanvas().
		WriteAt(xy.Origin, "aBcDe")

	// Act
	cell, ok := canvas.GetCell(xy.Origin.Right(2))

	// Assert
	Assert(t).That(ok, is.True, "Cell at this position")
	Assert(t).That(cell.Char(), is.EqualTo('c'), "Correct cell at this position")
}

func Test_Canvas_SetCell(t *testing.T) {
	// Arrange
	char := 'a'
	position := xy.NewPosition(1, 2)
	canvas := display.NewCanvas()

	// Act
	canvas.SetCell(position, cells.NewCell(char, ansi.NoAttributes))
	cell, ok := canvas.GetCell(position)

	// Assert
	Assert(t).That(ok, is.True, "Cell at this position")
	Assert(t).That(cell.Char(), is.EqualTo(char), "Correct cell at this position")
}

func Test_Canvas_Intersections(t *testing.T) {
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
	canvas := display.NewCanvas().
		WriteAt(xy.NewPosition(1, 0), horLine).
		WriteAt(xy.NewPosition(1, 2), horLine).
		WriteAt(xy.NewPosition(1, 4), horLine).
		WriteAt(xy.NewPosition(0, 1), vertLine).
		WriteAt(xy.NewPosition(2, 1), vertLine).
		WriteAt(xy.NewPosition(4, 1), vertLine)

	// Assert
	assertTile(t, canvas.Render(rules),
		"AhahB",
		"v v v",
		"chvhd",
		"v v v",
		"ChbhD")
}
