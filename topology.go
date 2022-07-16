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
		return x, y
	}

	return -1, -1
}

func TranslateTorus(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x, y
	}

	return (x + dx) % n, (y + dy) % n
}

func TranslateKleinBottle(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x, y
	}

	ny := y
	if (y+dy) < 0 || (y+dy) >= n {
		ny = (-1 * (y + dy)) % n
	}
	return (x + dx) % n, ny
}

func TranslateProjectivePlane(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x, y
	}

	nx := x
	ny := y
	if (y+dy) < 0 || (y+dy) >= n {
		ny = (-1 * (y + dy)) % n
	}
	if (x+dx) < 0 || (x+dx) >= n {
		nx = (-1 * (x + dx)) % n
	}
	return nx, ny
}

func TranslateSphere(x, y int, dx, dy int, n int) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x, y
	}
	// For boundary cases, points are reflected across straight line between (0,1) and (1, 0)
	return (-1 * (y + dy)) % n, (-1 * (x + dx)) % n
}
