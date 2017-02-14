package gridt

import (
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestFromBidimensional(t *testing.T) {
	t.Run("With zero items and unlimited size", func(t *testing.T) {
		ws, f := FromBidimensional([]string{}, Unlimited)
		const msg = "\nShould return an empty list of widths that fits: "
		if len(ws) != 0 || !f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})
	t.Run("With zero items and limited size", func(t *testing.T) {
		ws, f := FromBidimensional([]string{}, 10)
		const msg = "\nShould return an empty list of widths that fits: "
		if len(ws) != 0 || !f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})
	t.Run("With one item and unlimited size", func(t *testing.T) {
		ws, f := FromBidimensional([]string{"1234567890"}, Unlimited)
		const msg = "\nShould return a list with one width that fits: "
		if len(ws) != 1 || !f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})
	t.Run("With one item and sufficient size", func(t *testing.T) {
		ws, f := FromBidimensional([]string{"1234567890"}, 20)
		const msg = "\nShould return a list with one width that fits: "
		if len(ws) != 1 || !f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})
	t.Run("With one item and unsufficient size", func(t *testing.T) {
		ws, f := FromBidimensional([]string{"1234567890"}, 5)
		const msg = "\nShould return an empty list that does not fit: "
		if len(ws) != 0 || f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})
}
