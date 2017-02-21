package gridt

import (
	"bytes"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
)

// Dimensions contains the dimensions of the grid and the possibility to format it.
type Dimensions struct {
	ws []uint
	l  uint
	g  Grid
}

// Widths represents the widths of each column in the grid.
func (d Dimensions) Widths() []uint {
	return d.ws
}

// Lines represents the quantity of lines the grid has.
func (d Dimensions) Lines() uint {
	return d.l
}

// Columns represents the quantity of columns the grid has.
func (d Dimensions) Columns() uint {
	return uint(len(d.ws))
}

// String formats the grid into a string.
func (d Dimensions) String() string {
	var buf bytes.Buffer
	lines, columns := int(d.l), len(d.ws)
	for line := 0; line < lines; line++ {
		for column := 0; column < columns; column++ {
			var i int
			switch d.g.d {
			case LeftToRight:
				i = column + line*columns
			case TopToBottom:
				i = line + lines*column
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
