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
	co := uint(0)

	a, b, c, d := box[x-1][y-1], box[x-1][y+1], box[x+1][y-1], box[x+1][y+1]

	if a < 2 {
		co += a
	}
	if b < 2 {
		co += b
	}
	if c < 2 {
		co += c
	}
	if d < 2 {
		co += d
	}

	return co
}

func getMidpoint(box [][]uint) (uint, uint) {
	first := uint(len(box) / 2)
	if len(box) > 0 {
		return first, uint(len(box[0]) / 2)
	} else {
		return first, uint(0)
	}
}
