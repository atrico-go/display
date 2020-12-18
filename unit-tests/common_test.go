package unit_tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/atrico-go/display/common"
	"github.com/atrico-go/display/tile"

	. "github.com/atrico-go/core"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
	"github.com/atrico-go/testing/random"

	"github.com/atrico-go/display/xy"
)

var randomValues = random.NewValueGeneratorBuilder().
	WithDefaultStringLength(5).
	Build()

func assertTile(t *testing.T, element tile.Tile, lines ...string) {
	DisplayMultiline(element)
	width := common.GetMaxWidth(lines)
	// Size
	size := xy.NewSize(width, len(lines))
	Assert(t).That(element.Size(), is.EqualTo(size), "Size")
	paddedLines := padLines(lines, width)
	// Cells
	for y := 0; y < size.Y(); y++ {
		for x := 0; x < size.X(); x++ {
			pos := xy.NewPosition(x, y)
			cell, ok := element.GetCell(pos)
			Assert(t).That(ok, is.True, "Cell %v exists", pos)
			Assert(t).That(cell.Char(), is.EqualTo([]rune(paddedLines[y])[x]), "Cell %v contents", pos)
		}
	}
	// StringerMl
	mlString := element.StringMl()
	Assert(t).That(len(mlString), is.EqualTo(size.Y()), "Lines")
	for i := 0; i < size.Y(); i++ {
		Assert(t).That(mlString[i], is.EqualTo(paddedLines[i]), fmt.Sprintf("Line %d", i))
	}
	// Renderable
	rendered := element.Render()
	Assert(t).That(rendered, is.DeepEqualTo(element), "Renders to same")
}

func assertRenderable(t *testing.T, renderable tile.Renderable, lines ...string) {
	assertTile(t, renderable.Render(), lines...)
}

func padLines(lines []string, width int) []string {
	paddedLines := make([]string, len(lines))
	for i, line := range lines {
		paddedLines[i] = fmt.Sprintf("%s%s", line, strings.Repeat(" ", width-common.StringWidth(line)))
	}
	return paddedLines
}

func createTestElement(size xy.Size) (element tile.Tile, lines []string) {
	lines = make([]string, size.Y())
	for y := range lines {
		line := strings.Builder{}
		for x := 0; x < size.X(); x++ {
			line.WriteString(randomValues.StringOfLen(1))
		}
		lines[y] = line.String()
	}
	return tile.NewTile(lines), lines
}
