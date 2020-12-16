package unit_tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/atrico-go/display/tile"

	"github.com/atrico-go/display"
	"github.com/atrico-go/display/xy"
)

func Test_Border_None(t *testing.T) {
	// Arrange
	element, expectedLines := createTestElement(xy.NewSize(4, 4))

	// Act
	border := display.NewBorder(element, nil, nil, nil, nil)
	border2 := display.NewBorderBuilder().
		WithContent(element).
		Build()

	// Assert
	assertRenderable(t, border, expectedLines...)
	assertRenderable(t, border2, expectedLines...)
}

func Test_Border_Above(t *testing.T) {
	// Arrange
	width := 4
	char := '*'
	element, expectedLines := createTestElement(xy.NewSize(width, width))
	horizontal := strings.Repeat(string(char), width)

	// Act
	border := display.NewBorder(element, &char, nil, nil, nil)
	border2 := display.NewBorderBuilder().
		WithContent(element).
		WithAbove(char).
		Build()

	// Assert
	assertRenderable(t, border, append([]string{horizontal}, expectedLines...)...)
	assertRenderable(t, border2, append([]string{horizontal}, expectedLines...)...)
}

func Test_Border_Below(t *testing.T) {
	// Arrange
	width := 4
	char := '*'
	element, expectedLines := createTestElement(xy.NewSize(width, width))
	horizontal := strings.Repeat(string(char), width)

	// Act
	border := display.NewBorder(element, nil, &char, nil, nil)
	border2 := display.NewBorderBuilder().
		WithContent(element).
		WithBelow(char).
		Build()

	// Assert
	assertRenderable(t, border, append(expectedLines, horizontal)...)
	assertRenderable(t, border2, append(expectedLines, horizontal)...)
}

func Test_Border_Horizontal(t *testing.T) {
	// Arrange
	width := 4
	char := '*'
	element, expectedLines := createTestElement(xy.NewSize(width, width))
	horizontal := strings.Repeat(string(char), width)

	// Act
	border := display.NewBorder(element, &char, &char, nil, nil)
	border2 := display.NewBorderBuilder().
		WithContent(element).
		WithHorizontal(char).
		Build()

	// Assert
	assertRenderable(t, border, append([]string{horizontal}, append(expectedLines, horizontal)...)...)
	assertRenderable(t, border2, append([]string{horizontal}, append(expectedLines, horizontal)...)...)
}

func Test_Border_Left(t *testing.T) {
	// Arrange
	width := 4
	char := '*'
	element, expectedLines := createTestElement(xy.NewSize(width, width))

	// Act
	border := display.NewBorder(element, nil, nil, &char, nil)
	border2 := display.NewBorderBuilder().
		WithContent(element).
		WithLeft(char).
		Build()

	// Assert
	assertRenderable(t, border, prependToAll(expectedLines, char)...)
	assertRenderable(t, border2, prependToAll(expectedLines, char)...)
}

func Test_Border_Right(t *testing.T) {
	// Arrange
	width := 4
	char := '*'
	element, expectedLines := createTestElement(xy.NewSize(width, width))

	// Act
	border := display.NewBorder(element, nil, nil, nil, &char)
	border2 := display.NewBorderBuilder().
		WithContent(element).
		WithRight(char).
		Build()

	// Assert
	assertRenderable(t, border, appendToAll(expectedLines, char)...)
	assertRenderable(t, border2, appendToAll(expectedLines, char)...)
}

func Test_Border_Vertical(t *testing.T) {
	// Arrange
	width := 4
	char := '*'
	element, expectedLines := createTestElement(xy.NewSize(width, width))

	// Act
	border := display.NewBorder(element, nil, nil, &char, &char)
	border2 := display.NewBorderBuilder().
		WithContent(element).
		WithVertical(char).
		Build()

	// Assert
	assertRenderable(t, border, prependToAll(appendToAll(expectedLines, char), char)...)
	assertRenderable(t, border2, prependToAll(appendToAll(expectedLines, char), char)...)
}

func Test_Border_All(t *testing.T) {
	// Arrange
	width := 4
	char := '*'
	element, expectedLines := createTestElement(xy.NewSize(width, width))
	horizontal := fmt.Sprintf(" %s ", strings.Repeat(string(char), width))

	// Act
	border := display.NewBorder(element, &char, &char, &char, &char)
	border2 := display.NewBorderBuilder().
		WithContent(element).
		WithHorizontal(char).
		WithVertical(char).
		Build()

	// Assert
	assertRenderable(t, border, append([]string{horizontal}, append(prependToAll(appendToAll(expectedLines, char), char), horizontal)...)...)
	assertRenderable(t, border2, append([]string{horizontal}, append(prependToAll(appendToAll(expectedLines, char), char), horizontal)...)...)
}

func Test_Border_EmptyTileWithBorder(t *testing.T) {
	// Arrange
	tile := tile.NewSizedTile(xy.NewSize(3, 3), "")

	// Act
	border := display.NewBorderBuilder().
		WithContent(tile).
		WithHorizontal('*').
		WithVertical('*').
		Build()

	// Assert
	assertRenderable(t, border, " *** ", "*   *", "*   *", "*   *", " *** ")
}

func Test_Border_EmptyTileWithNoBorder(t *testing.T) {
	// Arrange
	tile := tile.NewSizedTile(xy.NewSize(3, 3), "")

	// Act
	border := display.NewBorderBuilder().
		WithContent(tile).
		Build()

	// Assert
	assertRenderable(t, border, "   ", "   ", "   ")
}

func Test_Border_Intersections(t *testing.T) {
	// Arrange
	tile := tile.NewSizedTile(xy.NewSize(3, 3), "")
	horizontal := 'h'
	vertical := 'v'
	rules := display.NewIntersectionRuleBuilder().
		AddIntersection(nil, &vertical, nil, &horizontal, 'A'). // top-left
		AddIntersection(nil, &vertical, &horizontal, nil, 'B'). // top-right
		AddIntersection(&vertical, nil, nil, &horizontal, 'C'). // bottom-left
		AddIntersection(&vertical, nil, &horizontal, nil, 'D'). // bottom-right
		Build()

	// Act
	border := display.NewBorderBuilder().
		WithContent(tile).
		WithHorizontal(horizontal).
		WithVertical(vertical).
		Build()

	// Assert
	assertTile(t, border.Render(rules),
		"AhhhB",
		"v   v",
		"v   v",
		"v   v",
		"ChhhD")
}

func prependToAll(lines []string, prefix rune) (newLines []string) {
	newLines = make([]string, len(lines))
	for i := range lines {
		newLines[i] = fmt.Sprintf("%s%s", string(prefix), lines[i])
	}
	return newLines
}

func appendToAll(lines []string, suffix rune) (newLines []string) {
	newLines = make([]string, len(lines))
	for i := range lines {
		newLines[i] = fmt.Sprintf("%s%s", lines[i], string(suffix))
	}
	return newLines
}
