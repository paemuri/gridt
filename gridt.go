package gridt

const (
	// LeftToRight is a direction in which the values will be written.
	LeftToRight direction = iota
	// TopToBottom is a direction in which the values will be written.
	TopToBottom
)

type direction int8

// Grid represents the values' grid, that will be exported as a pretty formatted string.
type Grid struct {
	v   []string
	d   direction
	sep string
}

// New returns a new Grid.
// `d` represents the direction in which the values will be written.
// `sep` represents the separator; a string that will be between each column.
func New(d direction, sep string) *Grid {
	return NewWithSize(d, sep, 0)
}

// NewWithSize returns a new Grid, with a pre-defined size of cells.
// `d` represents the direction in which the values will be written.
// `sep` represents the separator; a string that will be between each column.
// `size` is the pre-defined size of cells.
func NewWithSize(d direction, sep string, size uint) *Grid {
	return &Grid{make([]string, size), d, sep}
}

// Cells returns all cells of the grid.
func (g Grid) Cells() []string {
	return g.v
}

// Direction returns the direction in which the grid will be written.
func (g Grid) Direction() direction {
	return g.d
}

// Separator returns the separator; the string that will be between each column.
func (g Grid) Separator() string {
	return g.sep
}

// Add adds a cell to the grid.
// `s` is the value that will be added.
func (g *Grid) Add(s string) *Grid {
	g.v = append(g.v, s)
	return g
}

// AddRange adds a slice of cells to the grid.
// `s` is the values that will be added.
func (g *Grid) AddRange(s []string) *Grid {
	g.v = append(g.v, s...)
	return g
}

// Insert inserts a value in a specified position in the grid.
// `i` the position of the value.
// `s` is the value that will be added.
func (g *Grid) Insert(i uint, s string) *Grid {
	g.v = g.v[0 : len(g.v)+1]
	copy(g.v[i+1:], g.v[i:])
	g.v[i] = s
	return g
}

// Delete deletes a value in a specified position in the grid.
// `i` the position of the value.
func (g *Grid) Delete(i uint) *Grid {
	copy(g.v[i:], g.v[i+1:])
	g.v = g.v[0 : len(g.v)+1]
	return g
}

// FitIntoWidth formats the grid, based on a maximum width.
// `max` represents the maximum width of the grid, based on characters.
// `widths` represents the size, by characters, of each column. `len(widths)` works as the number of columns.
// `lines` represents the number of lines.
// `ok` says whether the the grid fits in the maximum width informed. If false, ignore the other values.
func (g Grid) FitIntoWidth(max uint) (widths []uint, lines uint, ok bool) {
	switch count := len(g.v); count {

	// If the slice is empty, returns empty grid that fits.
	case 0:
		return nil, 0, false

	// If it has one item, it is validated.
	case 1:
		if l := uint(len(g.v[0])); l <= max {
			return []uint{l}, 1, true
		}
		return nil, 0, false

	// If it has two or more items...
	default:

		// If the maximum size is zero, it is invalid.
		if max <= 0 {
			return nil, 0, false
		}

		// `lines` represents the minimum number of lines necessary.
		// This loop will check for every possibility.
		for lines := 1; lines <= count; lines++ {

			// `columns` represents the number of columns, based on the current number of lines.
			// It is the cells count, divided by the number of lines, rounded up.
			columns := count / lines
			if count%lines != 0 {
				columns++
			}

			// Calculates the free space...
			// Which is the maximum size, minus the total width of all the separators.
			// If there is no free space, this possibility is ignored.
			free := int(max) - ((columns - 1) * len(g.sep))
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
			if sum < free {
				return widths, uint(lines), true
			}
		}

		// If no possibility worked, than the cells does not fit the maximum size.
		return nil, 0, false
	}
}

// FitIntoColumns formats the grid, based on a maximum quantity of columns.
// `max` represents the maximum quantity of columns of the grid.
// `widths` represents the size, by characters, of each column. `len(widths)` works as the number of columns.
// `lines` represents the number of lines.
// `ok` says whether the the grid fits in the maximum width informed. If false, ignore the other values.
func (g Grid) FitIntoColumns(max uint) (widths []uint, lines uint, ok bool) {
	// If the maximum size is zero, it is invalid.
	if max <= 0 {
		return nil, 0, false
	}

	// `lines` represents the number of lines.
	// It is the cells count, divided by the number of maximum columns, rounded up.
	m := int(max)
	l := len(g.v) / m
	if len(g.v)%m != 0 {
		l++
	}
	return g.biggerFromEachColumn(l, m), uint(l), true
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
		if l := uint(len(vv)); l > widths[i] {
			widths[i] = l
		}
	}
	return widths
}
