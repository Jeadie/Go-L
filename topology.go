package main

const (
	Bordered        string = "BORDERED"
	Torus                  = "TORUS"
	Sphere				   = "SPHERE"
	KleinBottle            = "KLEIN_BOTTLE"
	ProjectivePlane        = "PROJECTIVE_PLANE"
)

var ALLOWED_TOPOLOGIES = []string{Bordered, Torus, KleinBottle, ProjectivePlane, Sphere}
const DefaultTopology = Bordered

// Returns true iff x is an allowed topology.
func isValidTopology(x string) bool {
	for _, t := range ALLOWED_TOPOLOGIES {
		if t == x {
			return true
		}
	}
	return false
}

// Describes how co-ordinates get transformed onn the boundary conditions
type TopologyTransformation func(x, y int, dx, dy int, n int) (int, int)

// GetTransformation returns the transformation function for a given topology.
func GetTransformation(topology string) TopologyTransformation {
	switch topology {
	case Bordered:
		return TranslateBordered
	case Torus:
		return TranslateTorus
	case KleinBottle:
		return TranslateKleinBottle
	case ProjectivePlane:
		return TranslateProjectivePlane
	case Sphere:
		return TranslateSphere
	default:
		return GetTransformation(DefaultTopology)
	}
}

func TranslateBordered(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x+dx, y+dy
	}

	return -1, -1
}

func TranslateTorus(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x+dx, y+dy
	}

	return mod(x + dx, n), mod(y + dy, n)
}

func TranslateKleinBottle(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x+dx, y+dy
	}

	ny := y
	if (y+dy) < 0 || (y+dy) >= n {
		ny = mod(-1 * (y + dy), n)
	}
	return mod(x + dx, n), ny
}

func TranslateProjectivePlane(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x+dx, y+dy
	}

	nx := x
	ny := y
	if (y+dy) < 0 || (y+dy) >= n {
		ny = mod(-1 * (y + dy), n)
	}
	if (x+dx) < 0 || (x+dx) >= n {
		nx = mod(-1 * (x + dx), n)
	}
	return nx, ny
}

func TranslateSphere(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x+dx, y+dy
	}

	//
	// on border
	if x+dx == n || y+dy == n {
		return mod(y+dy, n), mod(x+dx, n)
	}

	// For boundary cases, points are reflected across straight line between (0,n-1) and (n-1, 0)
	x, y = reflectSphere(x+dx, y+dy, n)

	// on border
	if x == n || y == n {
		return mod(y, n), mod(x, n)
	}

	// Now reflect back into square.
	x, y = reflectSquare(x, y, n)
	// on border
	if x == n || y == n {
		return mod(y, n), mod(x, n)
	}
	return x, y
}

// Reflect the point (x,y) back into the square bound by (0, 0), (n, n)
func reflectSquare(x,y, n int) (int, int) {
	return reflectSquareCoordinate(x, n), reflectSquareCoordinate(y, n)
}

func reflectSquareCoordinate(a, n int) int {
	if a < 0 {
		return -a
	} else if a > n {
		return n - (a-n)
	}
	return a
}

// Reflect  point (x,y) across y=n-x
func reflectSphere(x, y, n int) (int, int) {
	return n-y, n-x
}

// mod operation that handles negative number properly.
func mod(d, m int) int {
	var res = d % m

	if res < 0 {
		return res + m
	}
	return res
}
