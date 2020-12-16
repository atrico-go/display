package unit_tests

import (
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/display/xy"
)

func Test_Position_AllPositionsOrigin(t *testing.T) {
	// Arrange
	origin := xy.Origin
	size := xy.NewSize(3, 3)
	expected := []xy.Position{
		xy.NewPosition(0, 0), xy.NewPosition(1, 0), xy.NewPosition(2, 0),
		xy.NewPosition(0, 1), xy.NewPosition(1, 1), xy.NewPosition(2, 1),
		xy.NewPosition(0, 2), xy.NewPosition(1, 2), xy.NewPosition(2, 2),
	}

	// Act
	all := xy.AllPositions(origin, size)

	// Assert
	Assert(t).That(all, is.DeepEqualTo(expected), "correct values")
}

func Test_Position_AllPositionsPositiveOffset(t *testing.T) {
	// Arrange
	origin := xy.NewPosition(1, 2)
	size := xy.NewSize(3, 3)
	expected := []xy.Position{
		xy.NewPosition(1, 2), xy.NewPosition(2, 2), xy.NewPosition(3, 2),
		xy.NewPosition(1, 3), xy.NewPosition(2, 3), xy.NewPosition(3, 3),
		xy.NewPosition(1, 4), xy.NewPosition(2, 4), xy.NewPosition(3, 4),
	}

	// Act
	all := xy.AllPositions(origin, size)

	// Assert
	Assert(t).That(all, is.DeepEqualTo(expected), "correct values")
}

func Test_Position_AllPositionsNegativeOffset(t *testing.T) {
	// Arrange
	origin := xy.NewPosition(-2, -1)
	size := xy.NewSize(3, 3)
	expected := []xy.Position{
		xy.NewPosition(-2, -1), xy.NewPosition(-1, -1), xy.NewPosition(0, -1),
		xy.NewPosition(-2, 0), xy.NewPosition(-1, 0), xy.NewPosition(0, 0),
		xy.NewPosition(-2, 1), xy.NewPosition(-1, 1), xy.NewPosition(0, 1),
	}

	// Act
	all := xy.AllPositions(origin, size)

	// Assert
	Assert(t).That(all, is.DeepEqualTo(expected), "correct values")
}
