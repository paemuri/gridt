package gridt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	logMsgf   = "\n%s: %s"
	fatalMsgf = logMsgf + "\ncolumns = %v\nlines = %v\nfit = %v"
)

var (
	randomLists = [][]string{
		{"a_value", "another_value", "lots_of_cells", "small_value", "biiiiiiiiiiiig_vaaaaaaaaaalue", "with spaces, it's better to read", "abc", "123", "baby_u_n_me", "bla bla blablablablab", "I see the endings, now", "nothing is broken"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "b", "c"},
		{"a", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "c"},
		{"a", "b", "ccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"},
		{"aaaaa", "b", "c", "d", "e", "f", "g", "h", "i"},
	}
)

func fixedList(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = "1234567890"
	}
	return s
}

func TestFitIntoWidth(t *testing.T) {
	gridTestFitInto(t, func(g *Grid, c, w int) (Dimensions, bool) {
		return g.FitIntoWidth(w)
	})

}
func TestFitIntoColumns(t *testing.T) {
	gridTestFitInto(t, func(g *Grid, c, w int) (Dimensions, bool) {
		return g.FitIntoColumns(c)
	})
}

func gridTestFitInto(t *testing.T, functionToTest func(g *Grid, c, w int) (Dimensions, bool)) {
	for _, testCase := range []struct {
		values    []string
		maxWidth  int
		direction Direction
		separator string
		columns   int
		lines     int
		fits      bool
	}{
		// Empty lists.
		{fixedList(0), 10, TopToBottom, " ", 0, 0, false},
		{fixedList(0), 10, LeftToRight, " ", 0, 0, false},

		// Invalid maximum size.
		{fixedList(10), 0, TopToBottom, " ", 0, 0, false},
		{fixedList(10), 0, LeftToRight, " ", 0, 0, false},
		{fixedList(10), -10, TopToBottom, " ", 0, 0, false},
		{fixedList(10), -10, LeftToRight, " ", 0, 0, false},

		// Lists with one cell.
		{fixedList(1), 20, TopToBottom, " ", 1, 1, true},
		{fixedList(1), 5, TopToBottom, " ", 0, 0, false},
		{fixedList(1), 20, LeftToRight, " ", 1, 1, true},
		{fixedList(1), 5, LeftToRight, " ", 0, 0, false},

		// Lists with two cells.
		{fixedList(2), 30, TopToBottom, " ", 2, 1, true},
		{fixedList(2), 15, TopToBottom, " ", 1, 2, true},
		{fixedList(2), 5, TopToBottom, " ", 0, 0, false},
		{fixedList(2), 30, LeftToRight, " ", 2, 1, true},
		{fixedList(2), 15, LeftToRight, " ", 1, 2, true},
		{fixedList(2), 5, LeftToRight, " ", 0, 0, false},

		// Lists with three cells.
		{fixedList(3), 50, TopToBottom, " ", 3, 1, true},
		{fixedList(3), 30, TopToBottom, " ", 2, 2, true},
		{fixedList(3), 15, TopToBottom, " ", 1, 3, true},
		{fixedList(3), 5, TopToBottom, " ", 0, 0, false},
		{fixedList(3), 50, LeftToRight, " ", 3, 1, true},
		{fixedList(3), 30, LeftToRight, " ", 2, 2, true},
		{fixedList(3), 15, LeftToRight, " ", 1, 3, true},
		{fixedList(3), 5, LeftToRight, " ", 0, 0, false},

		// Random-sized lists.
		{randomLists[0], 72, TopToBottom, "  ", 3, 4, true},
		{randomLists[0], 72, LeftToRight, "  ", 2, 6, true},
		{randomLists[1], 100, TopToBottom, "^-._,^-._,^-._,^-._,^-._,^-._,^", 1, 3, true},
		{randomLists[1], 100, LeftToRight, "^-._,^-._,^-._,^-._,^-._,^-._,^", 1, 3, true},
		{randomLists[2], 100, TopToBottom, "^-._,^-._,^-._,^-._,^-._,^-._,^", 1, 3, true},
		{randomLists[2], 100, LeftToRight, "^-._,^-._,^-._,^-._,^-._,^-._,^", 1, 3, true},
		{randomLists[3], 100, TopToBottom, "^-._,^-._,^-._,^-._,^-._,^-._,^", 1, 3, true},
		{randomLists[3], 100, LeftToRight, "^-._,^-._,^-._,^-._,^-._,^-._,^", 1, 3, true},
		{randomLists[4], 9, TopToBottom, " ", 3, 3, true},
		{randomLists[4], 9, LeftToRight, " ", 3, 3, true},
	} {

		gridNew := New(testCase.direction, testCase.separator, testCase.values...)
		gridNewAdd := New(testCase.direction, testCase.separator).Add(testCase.values...)
		gridNewDeleteInsert := New(testCase.direction, testCase.separator, testCase.values...)
		for i, v := range testCase.values {
			gridNewDeleteInsert.Delete(i)
			gridNewDeleteInsert.Insert(i, v)
		}

		for _, g := range []*Grid{
			gridNew,
			gridNewAdd,
			gridNewDeleteInsert,
		} {
			d, f := functionToTest(g, testCase.columns, testCase.maxWidth)
			assert.Equal(t, testCase.values, g.Cells())
			assert.Equal(t, testCase.direction, g.Direction())
			assert.Equal(t, testCase.separator, g.Separator())
			assert.Equal(t, testCase.lines, d.l)
			assert.Equal(t, testCase.fits, f)
		}
	}
}
