package gridt

import (
	"fmt"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"

	logMsgf   = "\n%s: %s"
	fatalMsgf = logMsgf + "\ncolumns = %v\nlines = %v\nfit = %v"
)

var (
	randomLists = [][]string{
		{"a_value", "another_value", "lots_of_cells", "small_value", "biiiiiiiiiiiig_vaaaaaaaaaalue", "with spaces, it's better to read", "abc", "123", "baby_u_n_me", "bla bla blablablablab", "I see the endings, now", "nothing is broken"},
		{"a", "b", "ccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"},
	}
)

func fixedList(n int) []string {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, "1234567890")
	}
	return s
}

func fmtMsg(len int, f bool) string {
	var list, fit string
	if !f {
		list = "n empty list"
		fit = "does not fit"
	} else {
		fit = "fits"
		list = fmt.Sprintf(" list with %d column(s)", len)
	}
	return fmt.Sprintf("Should return a%s that %s", list, fit)
}

func TestFitIntoWidth(t *testing.T) {
	for _, c := range []struct {
		v []string  //=> cells' values
		m uint      //=> maximum width
		d Direction //=> direction
		s string    //=> separator
		c int       //=> columns
		l uint      //=> lines
		f bool      //=> if it fits
	}{
		// Empty lists.
		{fixedList(0), 10, TopToBottom, " ", 0, 0, false},
		{fixedList(0), 10, LeftToRight, " ", 0, 0, false},

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
	} {
		msg := fmtMsg(c.c, c.f)
		ws, l, f := NewWithCells(c.d, c.s, c.v...).FitIntoWidth(c.m)
		if len(ws) != c.c || l != c.l || f != c.f {
			t.Fatalf(fatalMsgf, msg, ballotX, ws, l, f)
		}
		t.Logf(logMsgf, msg, checkMark)
	}
}
