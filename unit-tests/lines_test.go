package unit_tests

import (
	"testing"

	"github.com/atrico-go/display"
)

func Test_Lines_Horizontal(t *testing.T) {
	// Arrange

	// Act
	line := display.NewHorizontalLine('*', 5)

	// Assert
	assertRenderable(t, line, "*****")
}

func Test_Lines_Vertical(t *testing.T) {
	// Arrange

	// Act
	line := display.NewVerticalLine('*', 4)

	// Assert
	assertRenderable(t, line, "*", "*", "*", "*")
}
