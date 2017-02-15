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

func TestFromBidimensional(t *testing.T) {
	for _, c := range []struct {
		v     []string
		m     uint
		d     direction
		s     string
		lenWs int
		l     uint
		f     bool
	}{
		// Empty lists.
		{[]string{}, 10, TopToBottom, " ", 0, 0, false},

		// Lists with one cell.
		{[]string{"1234567890"}, 20, TopToBottom, " ", 1, 1, true},
		{[]string{"1234567890"}, 5, TopToBottom, " ", 0, 0, false},

		// Lists with two cells.
		{[]string{"1234567890", "1234567890"}, 30, TopToBottom, " ", 2, 1, true},
		{[]string{"1234567890", "1234567890"}, 15, TopToBottom, " ", 1, 2, true},
		{[]string{"1234567890", "1234567890"}, 5, TopToBottom, " ", 0, 0, false},

		// Lists with three cells.
		{[]string{"1234567890", "1234567890", "1234567890"}, 50, TopToBottom, " ", 3, 1, true},
		{[]string{"1234567890", "1234567890", "1234567890"}, 30, TopToBottom, " ", 2, 2, true},
		{[]string{"1234567890", "1234567890", "1234567890"}, 15, TopToBottom, " ", 1, 3, true},
		{[]string{"1234567890", "1234567890", "1234567890"}, 5, TopToBottom, " ", 0, 0, false},
	} {
		msg := fmtMsg(c.lenWs, c.f)
		ws, l, f := FromBidimensional(c.v, c.m, c.d, c.s)
		if len(ws) != c.lenWs || l != c.l || f != c.f {
			t.Fatalf(fatalMsgf, msg, ballotX, ws, l, f)
		}
		t.Logf(logMsgf, msg, checkMark)
	}
}
