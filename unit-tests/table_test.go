package unit_tests

import (
	"testing"

	"github.com/atrico-go/display/tile"

	"github.com/atrico-go/display"
	"github.com/atrico-go/display/xy"
)

func Test_Table_Empty(t *testing.T) {
	// Arrange

	// Act
	table := display.NewTableBuilder().
		Build()

	// Assert
	assertRenderable(t, table)
}

func Test_Table_SetCells(t *testing.T) {
	// Arrange

	// Act
	table := display.NewTableBuilder().
		SetCell(0, 0, "a").
		SetCell(1, 1, "b").
		SetCell(2, 2, "c").
		Build()

	// Assert
	assertRenderable(t, table,
		"a",
		" b",
		"  c")
}

func Test_Table_SetCellsDifferentWidths(t *testing.T) {
	// Arrange

	// Act
	table := display.NewTableBuilder().
		SetCell(0, 0, "one").SetCell(1, 0, "-").
		SetCell(0, 1, "2").SetCell(1, 1, "-").
		SetCell(0, 2, "three").SetCell(1, 2, "-").
		Build()

	// Assert
	assertRenderable(t, table,
		"one  -",
		"2    -",
		"three-")
}

func Test_Table_SetCellsDifferentHeights(t *testing.T) {
	// Arrange

	// Act
	table := display.NewTableBuilder().
		SetCell(0, 0, []string{"2", "2"}).SetCell(0, 1, "-").
		SetCell(1, 0, "1").SetCell(1, 1, "-").
		SetCell(2, 0, []string{"3", "3", "3"}).SetCell(2, 1, "-").
		Build()

	// Assert
	assertRenderable(t, table,
		"213",
		"2 3",
		"  3",
		"---")
}

func Test_Table_EmptyTile(t *testing.T) {
	// Arrange

	// Act
	table := display.NewTableBuilder().
		SetCell(0, 0, []string{"2", "2"}).SetCell(0, 1, "-").
		SetCell(1, 0, "1").SetCell(1, 1, "-").
		SetCell(2, 0, tile.NewSizedTile(xy.NewSize(1, 3), "")).SetCell(2, 1, "-").
		Build()

	// Assert
	assertRenderable(t, table,
		"21 ",
		"2  ",
		"   ",
		"---")
}
func Test_Table_AppendRow(t *testing.T) {
	// Arrange
	cell00 := []string{"2", "2"}
	cell10 := "1"
	cell01 := "three"
	cell11 := "--"

	// Act
	table := display.NewTableBuilder().
		AppendRow(cell00, cell10).
		AppendRow(cell01, cell11).
		Build()

	// Assert
	assertRenderable(t, table,
		"2    1",
		"2    ",
		"three--")
}

func Test_Table_HorizontalSeparator(t *testing.T) {
	// Arrange

	// Act
	table := createTestTable().
		WithHorizontalSeparator('*').
		Build()

	// Assert
	assertRenderable(t, table,
		"abc",
		"***",
		"def",
		"***",
		"ghi")
}

func Test_Table_VerticalSeparator(t *testing.T) {
	// Arrange

	// Act
	table := createTestTable().
		WithVerticalSeparator('*').
		Build()

	// Assert
	assertRenderable(t, table,
		"a*b*c",
		"d*e*f",
		"g*h*i")
}

func Test_Table_BothSeparators(t *testing.T) {
	// Arrange

	// Act
	table := createTestTable().
		WithHorizontalSeparator('-').
		WithVerticalSeparator('|').
		Build()

	// Assert
	assertRenderable(t, table,
		"a|b|c",
		"- - -",
		"d|e|f",
		"- - -",
		"g|h|i")
}

func Test_Table_Intersections(t *testing.T) {
	// Arrange
	horizontal := '-'
	vertical := '|'
	rules := display.NewIntersectionRuleBuilder().
		AddIntersection(&vertical, &vertical, &horizontal, &horizontal, '+').
		Build()

	// Act
	table := createTestTable().
		WithHorizontalSeparator(horizontal).
		WithVerticalSeparator(vertical).
		Build()

	// Assert
	assertTile(t, table.Render(rules),
		"a|b|c",
		"-+-+-",
		"d|e|f",
		"-+-+-",
		"g|h|i")
}

func createTestTable() display.TableBuilder {
	return display.NewTableBuilder().
		AppendRow("a", "b", "c").
		AppendRow("d", "e", "f").
		AppendRow("g", "h", "i")

}
