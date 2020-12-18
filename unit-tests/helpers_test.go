package unit_tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/atrico-go/console/ansi"
	"github.com/atrico-go/console/ansi/color"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/display/common"
)

func Test_Helpers_MaxWidth(t *testing.T) {
	// Arrange
	lines := make([]string, 3)
	lines[0] = strings.Repeat("a", 2)
	lines[1] = strings.Repeat("b", 8)
	lines[2] = strings.Repeat("c", 3)

	// Act
	max := common.GetMaxWidth(lines)

	// Assert
	Assert(t).That(max, is.EqualTo(8), "Max width")
}

func Test_Helpers_StringWidthNormal(t *testing.T) {
	// Arrange
	str := randomValues.String()
	fmt.Println(str)

	// Act
	length := common.StringWidth(str)

	// Assert
	Assert(t).That(length, is.EqualTo(len(str)), "Correct length")
}

func Test_Helpers_StringWidthUnicode(t *testing.T) {
	// Arrange
	exp := 5
	runes := make([]rune, exp)
	for i := 0; i < exp; i++ {
		runes[i] = rune(randomValues.IntBetween(0x2500, 0x2580))
	}
	str := string(runes)
	fmt.Println(str)

	// Act
	length := common.StringWidth(str)

	// Assert
	Assert(t).That(length, is.EqualTo(exp), "Correct length")
}

func Test_Helpers_StringWidthColor(t *testing.T) {
	// Arrange
	raw := randomValues.String()
	attr := ansi.Attributes{Foreground: color.Red, Background: color.None}.SetThis()
	str := attr.ApplyTo(raw)
	fmt.Println(str)

	// Act
	length := common.StringWidth(str)

	// Assert
	Assert(t).That(length, is.EqualTo(len(raw)), "Correct length")
}
