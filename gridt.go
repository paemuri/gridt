package gridt

type direction int8

const (
	LeftToRight direction = iota
	TopToBottom

	Unlimited uint = 0
)

func FromBidimensional(values []string, max uint) (widths []int, ok bool) {
	switch len(values) {
	case 0:
		return []int{}, true
	case 1:
		firstLen := len(values[0])
		if max == Unlimited || int(max) >= firstLen {
			return []int{firstLen}, true
		}
		return nil, false
	default:
		return nil, false
	}
}
