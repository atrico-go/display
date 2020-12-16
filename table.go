package display

import (
	"github.com/atrico-go/display/common"
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

type TableBuilder interface {
	// Config
	WithHorizontalSeparator(separator rune) TableBuilder
	WithVerticalSeparator(separator rune) TableBuilder
	SetCell(x, y int, content interface{}) TableBuilder
	SetCellByPosition(pos xy.Position, content interface{}) TableBuilder
	AppendRow(content ...interface{}) TableBuilder
	// tile.Renderable
	Build() tile.Renderable
}

func NewTableBuilder() TableBuilder {
	return &table{cells: make(map[xy.Position]tile.Renderable)}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

type table struct {
	cells      map[xy.Position]tile.Renderable
	horizontal *rune
	vertical   *rune
}

func (t *table) WithHorizontalSeparator(separator rune) TableBuilder {
	t.horizontal = &separator
	return t
}

func (t *table) WithVerticalSeparator(separator rune) TableBuilder {
	t.vertical = &separator
	return t
}

func (t *table) SetCell(x, y int, content interface{}) TableBuilder {
	return t.SetCellByPosition(xy.NewPosition(x, y), content)
}

func (t *table) SetCellByPosition(pos xy.Position, content interface{}) TableBuilder {
	t.cells[pos] = tile.NewRenderable(content)
	return t
}

func (t *table) AppendRow(content ...interface{}) TableBuilder {
	// Find current extent
	y := t.countRowsAndColumns().Y()
	for x, cell := range content {
		t.SetCell(x, y, cell)
	}
	return t
}

func (t *table) Build() tile.Renderable {
	return *t
}

// Renderable
func (t table) Render(rules ...tile.RenderRule) tile.Tile {
	canvas := NewCanvas()
	cells, rows, cols := t.expandCells()
	yPos := 0
	for y := range cells {
		xPos := 0
		size := xy.NewSize(0, 0)
		for x := range cells[y] {
			cellB := NewBorderBuilder().
				WithContent(tile.NewSizedTile(xy.NewSize(cols[x], rows[y]), cells[y][x]))
			if y > 0 && t.horizontal != nil {
				cellB.WithAbove(*t.horizontal)
			}
			if x > 0 && t.vertical != nil {
				cellB.WithLeft(*t.vertical)
			}
			cell := cellB.Build().Render()
			size = cell.Size()
			canvas.WriteAt(xy.NewPosition(xPos, yPos), cell)
			xPos = xPos + size.X()
		}
		yPos = yPos + size.Y()
	}
	return canvas.Render(rules...)
}

// ----------------------------------------------------------------------------------------------------------------------------
// internal
// ----------------------------------------------------------------------------------------------------------------------------

// Cells in form [row(y)][col(x)]
// Rows/cols are slices width/height of each
func (t table) expandCells() (cells [][]tile.Tile, rows, cols []int) {
	size := t.countRowsAndColumns()
	cells = make([][]tile.Tile, size.Y())
	rows = make([]int, size.Y())
	cols = make([]int, size.X())
	for y := range cells {
		cells[y] = make([]tile.Tile, size.X())
		for x := range cells[y] {
			if elem, ok := t.cells[xy.NewPosition(x, y)]; ok {
				cells[y][x] = elem.Render()
			} else {
				cells[y][x] = tile.NewTile("")
			}
			rows[y] = common.MaxInt(rows[y], cells[y][x].Size().Y())
			cols[x] = common.MaxInt(cols[x], cells[y][x].Size().X())
		}
	}
	return cells, rows, cols
}

func (t table) countRowsAndColumns() xy.Size {
	coord := xy.NewPosition(-1, -1)
	for k, _ := range t.cells {
		coord = xy.UpperPosition(coord, k)
	}
	return xy.NewSize(coord.X()+1, coord.Y()+1)
}
