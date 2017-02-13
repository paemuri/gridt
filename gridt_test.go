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
		ws, f, err := FromBidimensional([]string{}, Unlimited)
		const msg = "\nShould return empty list of widths that fits, with no errors: "
		if len(ws) != 0 || !f || err != nil {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v\nerr = %v", msg, ballotX, ws, f, err)
		}
		t.Log(msg, checkMark)
	})
	t.Run("With zero items and limited size", func(t *testing.T) {
		ws, f, err := FromBidimensional([]string{}, 10)
		const msg = "\nShould return empty list of widths that fits, with no errors: "
		if len(ws) != 0 || !f || err != nil {
			t.Fatalf("%s%s\nwidths = %v\nfit = %v\nerr = %v", msg, ballotX, ws, f, err)
		}
		t.Log(msg, checkMark)
	})
}
