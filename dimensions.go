package gridt

import (
	"bytes"
	"strings"

	rw "github.com/mattn/go-runewidth"
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

			// Gets the position of the current line x column in the values list.
			var i int
			switch d.g.d {
			case LeftToRight:
				i = column + line*d.Columns()
			case TopToBottom:
				i = line + column*d.Lines()
			}

			// If `i` goes out of the grid, the loop is continued.
			if i >= d.Lines()*d.Columns() {
				continue
			}

			// If `i` represents a value...
			if i < len(d.g.v) {
				// Writes the cell value.
				cell := d.g.v[i]
				buf.WriteString(cell)
				// Fills the rest of the column with spaces.
				buf.WriteString(strings.Repeat(" ", d.ws[column]-rw.StringWidth(cell)))
			} else {
				// If does not, fills everything with spaces.
				buf.WriteString(strings.Repeat(" ", d.ws[column]))
			}

			// Writes the separator only if it is not the last column.
			if column+1 != d.Columns() {
				buf.WriteString(d.g.sep)
			}
		}

		// Finish a line.
		buf.WriteRune('\n')
	}
	return buf.String()
}
