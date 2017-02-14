package gridt

import (
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestFromBidimensional(t *testing.T) {
	// Empty list.
	t.Run("With zero items", func(t *testing.T) {
		ws, f := FromBidimensional([]string{}, 10)
		const msg = "\nShould return an empty list of widths that fits: "
		if len(ws) != 0 || !f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})

	// One item on the list.
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

	// Two items on the list.
	t.Run("With two items and sufficient size for both", func(t *testing.T) {
		ws, f := FromBidimensional([]string{"1234567890", "1234567890"}, 50)
		const msg = "\nShould return a list with two widths that fits: "
		if len(ws) != 2 || !f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})
	t.Run("With two items and sufficient size for one column", func(t *testing.T) {
		ws, f := FromBidimensional([]string{"1234567890", "1234567890"}, 15)
		const msg = "\nShould return a list with one width that fits: "
		if len(ws) != 1 || !f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})
	t.Run("With two items and unsufficient size for any", func(t *testing.T) {
		ws, f := FromBidimensional([]string{"1234567890", "1234567890"}, 5)
		const msg = "\nShould return an empty list that does not fit: "
		if len(ws) != 0 || f {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v", msg, ballotX, ws, f)
		}
		t.Log(msg, checkMark)
	})
}
