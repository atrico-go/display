package unit_tests

import (
	"testing"

	"github.com/atrico-go/display/tile"

	"github.com/atrico-go/display/xy"
)

func Test_Tile_CreateFromString(t *testing.T) {
	// Arrange
	content := "Hello world"

	// Act
	tile := tile.NewTile(content)

	// Assert
	assertTile(t, tile, content)
}

func Test_Tile_CreateFromPod(t *testing.T) {
	// Arrange
	content := 123

	// Act
	tile := tile.NewTile(content)

	// Assert
	assertTile(t, tile, "123")
}

func Test_Tile_CreateFromStringArray(t *testing.T) {
	// Arrange
	content := []string{"Hello", "world", "123"}

	// Act
	tile := tile.NewTile(content)

	// Assert
	assertTile(t, tile, content...)
}

type StringerObj struct {
	string
}

func (obj StringerObj) String() string {
	return obj.string
}

func Test_Tile_CreateFromStringerObject(t *testing.T) {
	// Arrange
	content := StringerObj{"Hello world"}

	// Act
	tile := tile.NewTile(content)

	// Assert
	assertTile(t, tile, content.String())
}

type StringerObj2 struct {
	string
}

func (obj StringerObj2) String() string {
	return obj.string
}

func Test_Tile_CreateFromObjectArray(t *testing.T) {
	// Arrange
	content := []interface{}{
		StringerObj{"Hello"},
		StringerObj2{"world"},
	}

	// Act
	tile := tile.NewTile(content)

	// Assert
	assertTile(t, tile, "Hello", "world")
}

func Test_Tile_CreateFromVariedArray(t *testing.T) {
	// Arrange
	content := []interface{}{
		StringerObj{"Hello"},
		"123",
		StringerObj2{"world"},
		456,
	}

	// Act
	tile := tile.NewTile(content)

	// Assert
	assertTile(t, tile, "Hello", "123", "world", "456")
}

type ElementObj struct {
	size     xy.Size
	contents []string
}

func (obj ElementObj) Size() xy.Size {
	return obj.size
}
func (obj ElementObj) StringMl(_ ...interface{}) []string {
	return obj.contents
}

func Test_Tile_CreateFromElement(t *testing.T) {
	// Arrange
	content := ElementObj{xy.NewSize(3, 3), []string{"abc", "def", "ghi"}}

	// Act
	tile := tile.NewTile(content)

	// Assert
	assertTile(t, tile, "abc", "def", "ghi")
}

type RenderableObj struct {
	contents []string
}

func (obj RenderableObj) Render(rules ...tile.RenderRule) tile.Tile {
	return tile.NewTile(obj.contents)
}

func Test_Tile_CreateFromRenderable(t *testing.T) {
	// Arrange
	content := RenderableObj{[]string{"abc", "def", "ghi"}}

	// Act
	tile := tile.NewTile(content)

	// Assert
	assertTile(t, tile, "abc", "def", "ghi")
}

func Test_Tile_SizedSame(t *testing.T) {
	// Arrange
	content := []string{"abc", "def", "ghi"}

	// Act
	tile := tile.NewSizedTile(xy.NewSize(3, 3), content)

	// Assert
	assertTile(t, tile, "abc", "def", "ghi")
}

func Test_Tile_SizedBigger(t *testing.T) {
	// Arrange
	content := []string{"abc", "def", "ghi"}

	// Act
	tile := tile.NewSizedTile(xy.NewSize(4, 4), content)

	// Assert
	assertTile(t, tile, "abc ", "def ", "ghi ", "    ")
}

func Test_Tile_SizedSmaller(t *testing.T) {
	// Arrange
	content := []string{"abc", "def", "ghi"}

	// Act
	tile := tile.NewSizedTile(xy.NewSize(2, 2), content)

	// Assert
	assertTile(t, tile, "ab", "de")
}
