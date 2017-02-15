package gridt

import (
	"fmt"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"

	logMsgf   = "\n%s: %s"
	fatalMsgf = logMsgf + "\nwidths = %v\nfit = %v"
)

func fmtMsg(len int, f bool) string {
	var list, fit string
	switch len {
	case 0:
		list = "n empty list of widths"
	case 1:
		list = " list with 1 width"
	default:
		list = fmt.Sprintf(" list with %d widths", len)
	}
	if f {
		fit = "fits"
	} else {
		fit = "does not fit"
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
		// Empty list.
		{[]string{}, 10, TopToBottom, " ", 0, 0, true},

		// List with one cell.
		{[]string{"1234567890"}, 20, TopToBottom, " ", 1, 1, true},
		{[]string{"1234567890"}, 5, TopToBottom, " ", 0, 0, false},

		// List with two cells.
		{[]string{"1234567890", "1234567890"}, 30, TopToBottom, " ", 2, 1, true},
		{[]string{"1234567890", "1234567890"}, 15, TopToBottom, " ", 1, 2, true},
		{[]string{"1234567890", "1234567890"}, 5, TopToBottom, " ", 0, 0, false},

		// List with three cells.
		{[]string{"1234567890", "1234567890", "1234567890"}, 50, TopToBottom, " ", 3, 1, true},
		{[]string{"1234567890", "1234567890", "1234567890"}, 30, TopToBottom, " ", 2, 2, true},
		{[]string{"1234567890", "1234567890", "1234567890"}, 15, TopToBottom, " ", 1, 3, true},
		{[]string{"1234567890", "1234567890", "1234567890"}, 5, TopToBottom, " ", 0, 0, false},
	} {
		msg := fmtMsg(c.lenWs, c.f)
		ws, l, f := FromBidimensional(c.v, c.m, c.d, c.s)
		if len(ws) != c.lenWs || l != c.l || f != c.f {
			t.Fatalf(fatalMsgf, msg, ballotX, ws, f)
		}
		t.Logf(logMsgf, msg, checkMark)
	}
}
