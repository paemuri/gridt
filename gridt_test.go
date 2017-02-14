package gridt

import (
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"

	logMsgf   = "\n%s: %s"
	fatalMsgf = logMsgf + "\nwidths = %v\nfit = %v"
)

func end(passed bool, msg string, ws []uint, f bool, t *testing.T) {
	if !passed {
		t.Fatalf(fatalMsgf, msg, ballotX, ws, f)
	}
	t.Logf(logMsgf, msg, checkMark)
}

func TestFromBidimensional(t *testing.T) {
	var msg string
	var passed, f bool
	var l uint
	var ws []uint

	t.Run("EmptyList", func(t *testing.T) {
		msg = "Should return an empty list of widths that fits"
		ws, l, f = FromBidimensional([]string{}, 10, TopToBottom, " ")
		passed = len(ws) == 0 && l == 0 && f
		end(passed, msg, ws, f, t)
	})

	t.Run("OneItem", func(t *testing.T) {
		t.Run("SufficientSize", func(t *testing.T) {
			msg = "Should return a list with one width that fits"
			ws, l, f = FromBidimensional([]string{"1234567890"}, 20, TopToBottom, " ")
			passed = len(ws) == 1 && l == 1 && f
			end(passed, msg, ws, f, t)
		})
		t.Run("UnsufficientSize", func(t *testing.T) {
			msg = "Should return an empty list that does not fit"
			ws, l, f = FromBidimensional([]string{"1234567890"}, 5, TopToBottom, " ")
			passed = len(ws) == 0 && l == 0 && !f
			end(passed, msg, ws, f, t)
		})
	})

	t.Run("TwoItems", func(t *testing.T) {
		t.Run("SufficientSizeForTwoColumns", func(t *testing.T) {
			msg = "Should return a list with two widths that fits"
			ws, l, f = FromBidimensional([]string{"1234567890", "1234567890"}, 50, TopToBottom, " ")
			passed = len(ws) == 2 && l == 1 && f
			end(passed, msg, ws, f, t)
		})
		t.Run("SufficientSizeForOneColumn", func(t *testing.T) {
			msg = "Should return a list with one width that fits"
			ws, l, f = FromBidimensional([]string{"1234567890", "1234567890"}, 15, TopToBottom, " ")
			passed = len(ws) == 1 && l == 2 && f
			end(passed, msg, ws, f, t)
		})
		t.Run("UnsufficientSizeFor", func(t *testing.T) {
			msg = "Should return an empty list that does not fit"
			ws, l, f = FromBidimensional([]string{"1234567890", "1234567890"}, 5, TopToBottom, " ")
			passed = len(ws) == 0 && l == 0 && !f
			end(passed, msg, ws, f, t)
		})
	})
}
