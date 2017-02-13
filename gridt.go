package gridt

type direction int8

const (
	LeftToRight direction = iota
	TopToBottom

	Unlimited uint = 0
)

func FromBidimensional(values []string, max uint) (widths []int, fit bool, err error) {
	return []int{}, true, nil
}
