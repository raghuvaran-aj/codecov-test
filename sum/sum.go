package sum

func AddIntegers(x int, y int) (int, error) {
	var z int
	if x != 0 && y != 0 {
		z = x + y
	}
	return z, nil
}
