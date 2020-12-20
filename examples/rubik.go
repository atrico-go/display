package main

import (
	"github.com/atrico-go/console/ansi"
	"github.com/atrico-go/console/ansi/color"
	"github.com/atrico-go/console/box_drawing"
	"github.com/atrico-go/core"

	"github.com/atrico-go/display"
	"github.com/atrico-go/display/rules"
	"github.com/atrico-go/display/tile"
	"github.com/atrico-go/display/xy"
)

func main() {
	front := createFace(color.Blue)
	back := createFace(color.Green)
	top := createFace(color.White)
	bottom := createFace(color.Yellow)
	left := createFace(color.Red)
	orange := ansi.Attributes{Foreground: color.Red, Background: color.Yellow}
	right := createFaceImpl(orange, "\u2592") // Orange
	size := front.Render().Size().Add(xy.NewSize(-1, -1))
	cube := display.NewPanel(
		display.PositionedRenderable{Position: xy.NewPosition(0, 0), Renderable: front},
		display.PositionedRenderable{Position: xy.NewPosition(0, -size.Y()), Renderable: top},
		display.PositionedRenderable{Position: xy.NewPosition(0, +size.Y()), Renderable: bottom},
		display.PositionedRenderable{Position: xy.NewPosition(-size.X(), 0), Renderable: left},
		display.PositionedRenderable{Position: xy.NewPosition(+size.X(), 0), Renderable: right},
		display.PositionedRenderable{Position: xy.NewPosition(+2*size.X(), 0), Renderable: back}).
		Render(rules.UnicodeIntersections)
	core.DisplayMultiline(cube)
}

func createFace(col color.Color) tile.Renderable {
	attr := ansi.Attributes{Foreground: color.None, Background: col}
	return createFaceImpl(attr, " ")
}

func createFaceImpl(attr ansi.Attributes, char string) tile.Renderable {
	tile := attr.SetThis().ApplyTo(char)
	table := display.NewTableBuilder().
		WithHorizontalSeparator(box_drawing.GetHorizontal(box_drawing.BoxSingle)).
		WithVerticalSeparator(box_drawing.GetVertical(box_drawing.BoxSingle)).
		AppendRow(tile, tile, tile).
		AppendRow(tile, tile, tile).
		AppendRow(tile, tile, tile).
		Build()
	borderH := box_drawing.GetHorizontal(box_drawing.BoxDouble)
	borderV := box_drawing.GetVertical(box_drawing.BoxDouble)
	return display.NewBorder(table, &borderH, &borderH, &borderV, &borderV)
}
