package gridt

import (
	"bytes"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
)

// Dimensions contains the dimensions of the grid and the possibility to format it.
type Dimensions struct {
	ws []int
	l  int
	g  Grid
}

// Widths represents the widths of each column in the grid.
func (d Dimensions) Widths() []int {
	return d.ws
}

// Lines represents the quantity of lines the grid has.
func (d Dimensions) Lines() int {
	return d.l
}

// Columns represents the quantity of columns the grid has.
func (d Dimensions) Columns() int {
	return len(d.ws)
}

// String formats the grid into a string.
func (d Dimensions) String() string {
	var buf bytes.Buffer
	for line := 0; line < d.Lines(); line++ {
		for column := 0; column < d.Columns(); column++ {
			var i int
			switch d.g.d {
			case LeftToRight:
				i = column + line*d.Columns()
			case TopToBottom:
				i = line + d.Lines()*column
			}

			if i >= len(d.g.v) {
				continue
			}

			cell := d.g.v[i]
			buf.WriteString(cell)
			buf.WriteString(strings.Repeat(" ", int(d.ws[column])-runewidth.StringWidth(cell)))
			buf.WriteString(d.g.sep)
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}
