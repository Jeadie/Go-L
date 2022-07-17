package gol

import "github.com/Jeadie/Go-L"

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

// CreateUpdateRule from a 32 bit update rule number. See README.md#2d-implementation for details.
func CreateUpdateRule(updateRuleNumber uint) main.UpdateRuleFn {
	// Only contain states that are alive, but will also map to true
	// So for alive states: isAlive, found := aliveStates[s] == true, true
	// So for dead states: isAlive, found := aliveStates[s] == false, false
	aliveStates := make(map[uint]bool)

	// Check bit i of updateRuleNumber to add alive states to map
	for i := uint(0); i < 32; i++ {
		if updateRuleNumber & (1 << i) > 0 {
			aliveStates[i] = true
		}
	}

	return func(box [][]uint) uint {
		state := 16 * box[0][1] + 8*box[1][1] + 4* box[2][1] +  2* box[1][0] + box[1][2]

		if a, b := aliveStates[state]; a && b {
			return 1
		}
		return 0
	}
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
