package gridt

const (
	LeftToRight direction = iota
	TopToBottom
)

type direction int8

func FromBidimensional(v []string, m uint, d direction, sep string) ([]uint, uint, bool) {
	count := len(v)
	switch count {

	// If the slice is empty, returns empty grid that fits.
	case 0:
		return []uint{}, 0, false

	// If it has one item, it is validated.
	case 1:
		l := uint(len(v[0]))
		if m >= l {
			return []uint{l}, 1, true
		}
		return nil, 0, false

	// If it has two or more items...
	default:

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
			free := int(m) - ((columns - 1) * len(sep))
			if free < 0 {
				continue
			}

			// Creates a slice of the widths of the columns.
			widths := make([]uint, columns)
			for i, vv := range v {
				// `v` represents the list of values.
				// `widths` represents the list of columns' widths.
				// `i` cannot be the index of the value on `v`, but its index on the line.
				// So, `i` is adjusted, based on the direction of the grid population.
				switch d {
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
