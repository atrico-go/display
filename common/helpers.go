package common

import "github.com/atrico-go/console/ansi"

func MinInt(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func MaxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func StringWidth(str string) int {
	width := 0
	for _, part := range ansi.ParseString(str) {
		width += len([]rune(part.String))
	}
	return width
}

func GetMaxWidth(lines []string) int {
	width := 0
	for _, line := range lines {
		width = MaxInt(width, StringWidth(line))
	}
	return width
}
