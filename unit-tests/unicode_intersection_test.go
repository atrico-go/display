package unit_tests

import (
	"testing"

	"github.com/atrico-go/console/box_drawing"

	"github.com/atrico-go/display"
	"github.com/atrico-go/display/rules"
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

func Test_UnicodeIntersections_Single(t *testing.T) {
	// Arrange

	// Act
	hor := box_drawing.GetHorizontal(box_drawing.BoxSingle)
	vert := box_drawing.GetVertical(box_drawing.BoxSingle)
	square := createSquare(hor, hor, hor, vert, vert, vert)

	// Assert
	assertTile(t, square.Render(rules.UnicodeIntersections),
		"┌─┬─┐",
		"│ │ │",
		"├─┼─┤",
		"│ │ │",
		"└─┴─┘")
}

func Test_UnicodeIntersections_Double(t *testing.T) {
	// Arrange

	// Act
	hor := box_drawing.GetHorizontal(box_drawing.BoxDouble)
	vert := box_drawing.GetVertical(box_drawing.BoxDouble)
	square := createSquare(hor, hor, hor, vert, vert, vert)

	// Assert
	assertTile(t, square.Render(rules.UnicodeIntersections),
		"╔═╦═╗",
		"║ ║ ║",
		"╠═╬═╣",
		"║ ║ ║",
		"╚═╩═╝")
}

func Test_UnicodeIntersections_Heavy(t *testing.T) {
	// Arrange

	// Act
	hor := box_drawing.GetHorizontal(box_drawing.BoxHeavy)
	vert := box_drawing.GetVertical(box_drawing.BoxHeavy)
	square := createSquare(hor, hor, hor, vert, vert, vert)

	// Assert
	assertTile(t, square.Render(rules.UnicodeIntersections),
		"┏━┳━┓",
		"┃ ┃ ┃",
		"┣━╋━┫",
		"┃ ┃ ┃",
		"┗━┻━┛")
}

func Test_UnicodeIntersections_DoubleSingle(t *testing.T) {
	// Arrange

	// Act
	hor := box_drawing.GetHorizontal(box_drawing.BoxDouble)
	vert := box_drawing.GetVertical(box_drawing.BoxSingle)
	square := createSquare(hor, hor, hor, vert, vert, vert)

	// Assert
	assertTile(t, square.Render(rules.UnicodeIntersections),
		"╒═╤═╕",
		"│ │ │",
		"╞═╪═╡",
		"│ │ │",
		"╘═╧═╛")
}

func Test_UnicodeIntersections_SingleDouble(t *testing.T) {
	// Arrange

	// Act
	hor := box_drawing.GetHorizontal(box_drawing.BoxSingle)
	vert := box_drawing.GetVertical(box_drawing.BoxDouble)
	square := createSquare(hor, hor, hor, vert, vert, vert)

	// Assert
	assertTile(t, square.Render(rules.UnicodeIntersections),
		"╓─╥─╖",
		"║ ║ ║",
		"╟─╫─╢",
		"║ ║ ║",
		"╙─╨─╜")
}

func Test_UnicodeIntersections_HeavySingle(t *testing.T) {
	// Arrange

	// Act
	hor := box_drawing.GetHorizontal(box_drawing.BoxHeavy)
	vert := box_drawing.GetVertical(box_drawing.BoxSingle)
	square := createSquare(hor, hor, hor, vert, vert, vert)

	// Assert
	assertTile(t, square.Render(rules.UnicodeIntersections),
		"┍━┯━┑",
		"│ │ │",
		"┝━┿━┥",
		"│ │ │",
		"┕━┷━┙")
}

func Test_UnicodeIntersections_SingleHeavy(t *testing.T) {
	// Arrange

	// Act
	hor := box_drawing.GetHorizontal(box_drawing.BoxSingle)
	vert := box_drawing.GetVertical(box_drawing.BoxHeavy)
	square := createSquare(hor, hor, hor, vert, vert, vert)

	// Assert
	assertTile(t, square.Render(rules.UnicodeIntersections),
		"┎─┰─┒",
		"┃ ┃ ┃",
		"┠─╂─┨",
		"┃ ┃ ┃",
		"┖─┸─┚")
}

func createSquare(top, equator, bottom rune, left, middle, right rune) tile.Renderable {
	horLineT := display.NewHorizontalLine(top, 3)
	horLineE := display.NewHorizontalLine(equator, 3)
	horLineB := display.NewHorizontalLine(bottom, 3)
	vertLineL := display.NewVerticalLine(left, 3)
	vertLineM := display.NewVerticalLine(middle, 3)
	vertLineR := display.NewVerticalLine(right, 3)
	return display.NewCanvas().
		WriteAt(xy.NewPosition(1, 0), horLineT).
		WriteAt(xy.NewPosition(1, 2), horLineE).
		WriteAt(xy.NewPosition(1, 4), horLineB).
		WriteAt(xy.NewPosition(0, 1), vertLineL).
		WriteAt(xy.NewPosition(2, 1), vertLineM).
		WriteAt(xy.NewPosition(4, 1), vertLineR)

}
