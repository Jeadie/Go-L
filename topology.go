package main

const DefaultTopology = Bordered

// TranslateVertex applies a differential change dx, dy to a point x, y and apply the boundary rules according to the topology
func TranslateVertex(x, y int, dx, dy int, n int, topology string) (int, int) {
	// Efficient non-boundary case return.
	if x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < n {
		return x, y
	}

	switch topology {
	case Bordered:
		return TranslateBordered(x, y, dx, dy, n)
	case Torus:
		return TranslateTorus(x, y, dx, dy, n)
	case KleinBottle:
		return TranslateKleinBottle(x, y, dx, dy, n)
	case ProjectivePlane:
		return TranslateProjectivePlane(x, y, dx, dy, n)
	case Sphere:
		return TranslateSphere(x, y, dx, dy, n)
	default:
		return TranslateVertex(x, y, dx, dy, n, DefaultTopology)
	}
}

func TranslateBordered(x, y int, dx, dy int, n int) (int, int) {
	return -1, -1
}

func TranslateTorus(x, y int, dx, dy int, n int) (int, int) {
	return (x + dx) % n, (y + dy) % n
}

func TranslateKleinBottle(x, y int, dx, dy int, n int) (int, int) {
	ny := y
	if (y+dy) < 0 || (y+dy) >= n {
		ny = (-1 * (y + dy)) % n
	}
	return (x + dx) % n, ny
}

func TranslateProjectivePlane(x, y int, dx, dy int, n int) (int, int) {
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
	// For boundary cases, points are reflected across straight line between (0,1) and (1, 0)
	return (-1 * (y + dy)) % n, (-1 * (x + dx)) % n
}
