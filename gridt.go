package gridt

const (
	LeftToRight direction = iota
	TopToBottom
)

type direction int8

func FromBidimensional(values []string, max uint) (widths []uint, ok bool) {
	switch len(values) {
	case 0:
		return []uint{}, true
	case 1:
		firstLen := uint(len(values[0]))
		if max >= firstLen {
			return []uint{firstLen}, true
		}
		return nil, false
	default:
		return nil, false
	}
}
