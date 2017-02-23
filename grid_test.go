package gridt

import (
	"fmt"
	"testing"
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

func fmtMsg(i, ii, len int, f bool) string {
	var list, fit string
	if !f {
		list = "n empty list"
		fit = "does not fit"
	} else {
		fit = "fits"
		list = fmt.Sprintf(" list with %d column(s)", len)
	}
	return fmt.Sprintf("#%d #%d Should return a%s that %s", i, ii, list, fit)
}

func TestFitIntoWidth(t *testing.T) {
	gridTestFitInto(t, func(g *Grid, c, w int) (Dimensions, bool) { return g.FitIntoWidth(w) })

}
func TestFitIntoColumns(t *testing.T) {
	gridTestFitInto(t, func(g *Grid, c, w int) (Dimensions, bool) { return g.FitIntoColumns(c) })
}

func gridTestFitInto(t *testing.T, test func(g *Grid, c, w int) (Dimensions, bool)) {
	for i, c := range []struct {
		v []string  //=> cells' values
		m int       //=> maximum width
		d Direction //=> direction
		s string    //=> separator
		c int       //=> columns
		l int       //=> lines
		f bool      //=> if it fits
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
		gNewWithCells := NewWithCells(c.d, c.s, c.v...)
		gNewAdd := New(c.d, c.s).Add(c.v...)
		gNewInsertDelete := NewWithCells(c.d, c.s, c.v...)
		for i, v := range c.v {
			gNewInsertDelete.Delete(i)
			gNewInsertDelete.Insert(i, v)
		}
		for ii, g := range []*Grid{
			gNewWithCells,
			gNewAdd,
			gNewInsertDelete,
		} {
			msg := fmtMsg(i, ii, c.c, c.f)
			d, f := test(g, c.c, c.m)
			cellsOk := len(c.v) == len(g.Cells())
			for i, cell := range g.Cells() {
				if c.v[i] != cell {
					cellsOk = false
					break
				}
			}
			if g.Direction() != c.d ||
				g.Separator() != c.s ||
				!cellsOk ||
				d.l != c.l ||
				f != c.f {
				t.Fatalf(fatalMsgf, msg, ballotX, d.ws, d.l, f)
			}
			t.Logf(logMsgf, msg, checkMark)
		}
	}
}
