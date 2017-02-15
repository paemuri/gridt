package gridt

func divUp(a, b int) int {
	c := a / b
	if a%b != 0 {
		c++
	}
	return c
}
