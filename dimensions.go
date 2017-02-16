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
	var buf bytes.Buffer
	switch d.g.d {
	case TopToBottom:
		last, counter := 0, -1
		linesBuf := make([]bytes.Buffer, d.l)
		for i, cell := range d.g.v {
			i /= int(d.l)
			if i != last {
				last = i
				counter = 0
			} else {
				counter++
			}
			linesBuf[counter].WriteString(cell)
			linesBuf[counter].WriteString(strings.Repeat(" ", int(d.ws[i])-len(cell)))
			if len(d.ws) != i+1 {
				linesBuf[counter].WriteString(d.g.sep)
			}
		}
		if counter != int(d.l) {
			for counter++; counter < int(d.l); counter++ {
				linesBuf[counter].WriteString(strings.Repeat(" ", int(d.ws[len(d.ws)-1])))
			}
		}
		for i, b := range linesBuf {
			buf.WriteString(b.String())
			if i+1 != len(linesBuf) {
				buf.WriteRune('\n')
			}
		}
	case LeftToRight:
		var i int
		var cell string
		for i, cell = range d.g.v {
			col := i % len(d.ws)
			buf.WriteString(cell)
			buf.WriteString(strings.Repeat(" ", int(d.ws[col])-len(cell)))
			if i+1 != len(d.g.v) {
				if len(d.ws) == col+1 {
					buf.WriteRune('\n')
				} else {
					buf.WriteString(d.g.sep)
				}
			}
		}
		if rest := i % len(d.ws); rest != 0 {
			for rest++; rest < len(d.ws); rest++ {
				buf.WriteString(d.g.sep)
				buf.WriteString(strings.Repeat(" ", int(d.ws[rest])))
			}
		}
	}
	return buf.String()
}
