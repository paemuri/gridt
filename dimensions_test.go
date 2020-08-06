package gridt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	for _, testCase := range []struct {
		values    []string
		direction Direction
		separator string
		widths    []int
		lines     int
	}{
		{[]string{"aaa", "bbb", "ccc"}, TopToBottom, "^-._,^-._,^-._,^-._,^-._,^", []int{3}, 3},
		{[]string{"aaa", "bbb", "ccc"}, LeftToRight, "^-._,^-._,^-._,^-._,^-._,^", []int{3}, 3},
		{[]string{"aaa", "bbb", "ccc", "aaa", "bbb", "ccc"}, TopToBottom, "---", []int{3, 3}, 3},
		{[]string{"aaa", "bbb", "ccc", "aaa", "bbb", "ccc"}, LeftToRight, "---", []int{3, 3}, 3},
		{[]string{"x", "xxxxxxxxxx", "xxx", "xxxxxxxxx", "xxxx", "xxxxxxx", "xxxxxx"}, TopToBottom, "-", []int{10, 7}, 4},
		{[]string{"x", "xxxxxxxxxx", "xxx", "xxxxxxxxx", "xxxx", "xxxxxxx", "xxxxxx"}, LeftToRight, "-", []int{6, 10}, 4},
	} {
		dimensions := Dimensions{testCase.widths, testCase.lines, Grid{testCase.values, testCase.direction, testCase.separator}}

		dimensionsSize := len(testCase.separator) * (len(testCase.widths) - 1)
		for _, w := range dimensions.Widths() {
			dimensionsSize += w
		}
		dimensionsSize++
		dimensionsSize *= testCase.lines

		assert.Equal(t, testCase.lines, dimensions.Lines())
		assert.Equal(t, len(testCase.widths), dimensions.Columns())
		assert.Equal(t, testCase.widths, dimensions.Widths())
		assert.Equal(t, len(dimensions.String()), dimensionsSize)
	}
}
