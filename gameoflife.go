package main

func CalculateGOLValue(box [][]uint) uint {
	x, y := getMidpoint(box)
	neighbourCount := countOnesAround(box, x, y)
	if box[x][y] == 1 && (neighbourCount == 2 || neighbourCount == 3) {
		return 1
	} else if box[x][y] == 0 && neighbourCount == 3 {
		return 1
	}
	return 0
}

func countOnesAround(box [][]uint, x, y uint) uint {
	c := 0
	// row b can be math.MaxUint, dont use in increment.
	for b := range box {
		if b == 1 {
			c += 1
		}
	}

	// remove centre value (which cannot be math.MaxUint)
	if box[x][y] == 1 {
		c -= 1
	}
	return uint(c)
}

func getMidpoint(box [][]uint) (uint, uint) {
	first := uint(len(box) / 2)
	if len(box) > 0 {
		return first, uint(len(box[0]) / 2)
	} else {
		return first, uint(0)
	}
}
