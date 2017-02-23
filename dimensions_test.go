package gridt

import "testing"

func TestString(t *testing.T) {
	for i, c := range []struct {
		v []string
		d Direction
		s string
		w []int
		l int
	}{
		{[]string{"aaa", "bbb", "ccc"}, TopToBottom, "^-._,^-._,^-._,^-._,^-._,^", []int{3}, 3},
		{[]string{"aaa", "bbb", "ccc"}, LeftToRight, "^-._,^-._,^-._,^-._,^-._,^", []int{3}, 3},
		{[]string{"aaa", "bbb", "ccc", "aaa", "bbb", "ccc"}, TopToBottom, "---", []int{3, 3}, 3},
		{[]string{"aaa", "bbb", "ccc", "aaa", "bbb", "ccc"}, LeftToRight, "---", []int{3, 3}, 3},
		{[]string{"x", "xxxxxxxxxx", "xxx", "xxxxxxxxx", "xxxx", "xxxxxxx", "xxxxxx"}, TopToBottom, "-", []int{10, 7}, 4},
		{[]string{"x", "xxxxxxxxxx", "xxx", "xxxxxxxxx", "xxxx", "xxxxxxx", "xxxxxx"}, LeftToRight, "-", []int{6, 10}, 4},
	} {
		t.Logf("#%d The string returned should correspond to the dimensions.", i)
		d := Dimensions{c.w, c.l, Grid{c.v, c.d, c.s}}
		s := d.String()
		widthsOk := len(d.Widths()) == len(c.w)
		var widthsSum int
		for i, w := range d.Widths() {
			if w != c.w[i] {
				widthsOk = false
				break
			}
			widthsSum += w
		}
		if d.Lines() != c.l ||
			d.Columns() != len(c.w) ||
			!widthsOk ||
			len(s) != (widthsSum+(len(c.s)*(len(c.w)-1))+1)*c.l {
			t.Log(len(s))
			t.Log((widthsSum + (len(c.s) * (len(c.w) - 1)) + 1) * c.l)
			t.Log(s)
			t.Fatalf(ballotX)
		}
		t.Log(checkMark)
	}
}
