package gridt

import (
	"bytes"
	"strings"
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
	var buffer bytes.Buffer
	for i, cell := range d.g.v {
		switch d.g.d {
		case TopToBottom:
			i /= int(d.l)
		case LeftToRight:
			i %= len(d.ws)
		}
		buffer.WriteString(cell)
		buffer.WriteString(strings.Repeat(" ", int(d.ws[i])-len(cell)))
		buffer.WriteString(d.g.sep)
		if len(d.ws) == i+1 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}
