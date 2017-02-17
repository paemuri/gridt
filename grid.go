package gridt

import (
	runewidth "github.com/mattn/go-runewidth"
)

const (
	// LeftToRight is a direction in which the values will be written.
	// It goes from the first cell (0,0) to the end of the line, returning to the beggining of the second line.
	// Exactly the same as a typewritter.
	LeftToRight Direction = iota
	// TopToBottom is a direction in which the values will be written.
	// It goes from the first cell (0,0) to the bottom of the column, returning to the top of the second column.
	// Exactly the same as how `ls` command works by default.
	TopToBottom
)

// Direction represents the direction in which the values will be written.
type Direction int8

// Grid represents the values' grid, that will be exported as a pretty formatted string.
type Grid struct {
	v   []string
	d   Direction
	sep string
}

// New returns a new Grid.
// `d` represents the direction in which the values will be written.
// `sep` represents the separator; a string that will be between each column.
func New(d Direction, sep string) *Grid {
	return &Grid{make([]string, 0), d, sep}
}

// NewWithCells returns a new Grid, with pre-defined cells.
// `d` represents the direction in which the values will be written.
// `sep` represents the separator; a string that will be between each column.
// `s` is the cells that will be added.
func NewWithCells(d Direction, sep string, s ...string) *Grid {
	return New(d, sep).Add(s...)
}

// Cells returns all cells of the grid.
func (g Grid) Cells() []string {
	return g.v
}

// Direction returns the direction in which the grid will be written.
func (g Grid) Direction() Direction {
	return g.d
}

// Separator returns the separator; the string that will be between each column.
func (g Grid) Separator() string {
	return g.sep
}

// Add adds a cell to the grid.
// `s` is the values that will be added.
func (g *Grid) Add(s ...string) *Grid {
	g.v = append(g.v, s...)
	return g
}

// Insert inserts a value in a specified position in the grid.
// `i` the position of the value.
// `s` is the value that will be added.
func (g *Grid) Insert(i uint, s ...string) *Grid {
	ii := int(i)
	g.v = g.v[0 : len(g.v)+len(s)]
	copy(g.v[ii+len(s):], g.v[ii:])
	for si, ss := range s {
		g.v[ii+si] = ss
	}
	return g
}

// Delete deletes a value in a specified position in the grid.
// `i` the position of the value.
func (g *Grid) Delete(i ...uint) *Grid {
	for ii := range i {
		copy(g.v[ii:], g.v[ii+1:])
	}
	g.v = g.v[0 : len(g.v)-len(i)]
	return g
}

// FitIntoWidth formats the grid, based on a maximum width.
// `max` represents the maximum width of the grid, based on characters.
// `dim` represents the dimensions of the grid, used for formatting. See `Dimensions`.
// `ok` says whether the the grid fits in the maximum width informed. If false, discard `dim`.
func (g Grid) FitIntoWidth(max uint) (dim Dimensions, ok bool) {
	switch count := len(g.v); count {

	// If the slice is empty, returns empty grid that fits.
	case 0:
		return Dimensions{}, false

	// If it has one item, it is validated.
	case 1:
		if l := uint(runewidth.StringWidth(g.v[0])); l <= max {
			return Dimensions{[]uint{l}, 1, g}, true
		}
		return Dimensions{}, false

	// If it has two or more items...
	default:

		// If the maximum size is zero, it is invalid.
		if max <= 0 {
			return Dimensions{}, false
		}

		// `lines` represents the minimum number of lines necessary.
		// This loop will check for every possibility.
		for lines := 1; lines <= count; lines++ {

			// `columns` represents the number of columns, based on the current number of lines.
			// It is the cells count, divided by the number of lines, rounded up.
			columns := divUp(count, lines)

			// Calculates the free space...
			// Which is the maximum size, minus the total width of all the separators.
			// If there is no free space, this possibility is ignored.
			free := int(max) - ((columns - 1) * runewidth.StringWidth(g.sep))
			if free < 0 {
				continue
			}

			// Gets bigger widths from each column. :P
			widths := g.biggerFromEachColumn(lines, columns)

			// If the sum of all widths fits the free space, then the possibility is reality!
			var sum int
			for _, width := range widths {
				sum += int(width)
			}
			if sum <= free {
				return Dimensions{widths, uint(lines), g}, true
			}
		}

		// If no possibility worked, than the cells does not fit the maximum size.
		return Dimensions{}, false
	}
}

// FitIntoColumns formats the grid, based on a maximum quantity of columns.
// `max` represents the maximum quantity of columns of the grid.
// `dim` represents the dimensions of the grid, used for formatting. See `Dimensions`.
// `ok` says whether the the grid fits in the maximum width informed. If false, discard `dim`.
func (g Grid) FitIntoColumns(max uint) (dim Dimensions, ok bool) {
	// If the maximum size is zero, it is invalid.
	if max <= 0 {
		return Dimensions{}, false
	}

	// `lines` represents the number of lines.
	// It is the cells count, divided by the number of maximum columns, rounded up.
	l := divUp(len(g.v), int(max))
	return Dimensions{g.biggerFromEachColumn(l, int(max)), uint(l), g}, true
}

func (g Grid) biggerFromEachColumn(lines, columns int) []uint {
	// Creates a slice of the widths of the columns.
	widths := make([]uint, columns)
	for i, vv := range g.v {
		// `v` represents the list of values.
		// `widths` represents the list of columns' widths.
		// `i` cannot be the index of the value on `v`, but its index on the line.
		// So, `i` is adjusted, based on the direction of the grid population.
		switch g.d {
		case TopToBottom:
			i /= lines
		case LeftToRight:
			i %= columns
		}

		// Now, `i` represents the index of the column (or cell on the line).
		// `widths[i]` is substituted by the current value, if the latter is bigger.
		// `widths[i]` represents the bigger value on the `i` column.
		if l := uint(runewidth.StringWidth(vv)); l > widths[i] {
			widths[i] = l
		}
	}
	return widths
}
