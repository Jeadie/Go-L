package gol

import (
	"constraints"
	"github.com/Jeadie/pkg/Go-L"
	"math"
	"math/rand"
)

// Node can be used in Lattice method signatures
type Node constraints.Ordered
type IntPair = [2]int


type Lattice[T Node] struct {
	grid       []T
	topologyFn TopologyTransformation
	n          uint
	null       T
	formatter  func(x T) string
	updateRule func([][]T) T
}

// GetValue on lattice at coordinate (x, y)
func (l *Lattice[T]) GetValue(x int, y int) T {
	nx, ny := l.topologyFn(x, y, 0,0, int(l.n))
	if nx == -1 || ny == -1 { return l.null }
	return l.grid[(nx * int(l.n)) + ny]
}

// SetValue on lattice at coordinate (x, y) to v
func (l *Lattice[T]) SetValue(x int, y int, v T)  {
	nx, ny := l.topologyFn(x, y, 0,0, int(l.n))
	l.grid[(nx*int(l.n))+ny] = v
}

// GetValuesAround returns all values around a coordinate within an L1 distance of w.
func (l *Lattice[T]) GetValuesAround(x int, y int, w int) [][]T {
	rows := make([][]T, 2*w+1)

	for i := 0; i < 2*w + 1; i++ {
		rows[i] = make([]T, 2*w + 1)
		for j := 0; j < 2*w+1; j++ {

			rows[i][j] = l.GetValue(x+i-w, y+j-w)
		}
	}
	return rows
}

// Copy a Lattice struct, and all its references, to a new instance (i.e. deep copy).
func (l *Lattice[T]) Copy() *Lattice[T] {
	readG :=  make([]T, len(l.grid))
	copy(readG, l.grid)
	return &Lattice[T]{
		grid:       readG,
		topologyFn: l.topologyFn,
		n:          l.n,
		updateRule: l.updateRule,
	}
}

// GetLatticeCoordinates into a channel of all coordinates on the lattice.
func (l *Lattice[T]) GetLatticeCoordinates() chan IntPair {
	out := make(chan IntPair)
	go func(out chan IntPair) {
		defer close(out)
		for i := 0; i < int(l.n); i++ {
			for j := 0; j < int(l.n); j++ {
				out <- [2]int{i, j}
			}
		}
	}(out)
	return out
}

// UpdatePair value within a lattice at a given co-ordinate. if `readL` is nil,
// will read from referenced Lattice. This allows for parallel update of
// lattice without conflict.
func (l *Lattice[T]) UpdatePair(i IntPair, readL *Lattice[T]) bool {
	x, y := i[0], i[1]
	var box [][]T
	if readL != nil {
		box = readL.GetValuesAround(x, y, 1)
	} else {
		box = l.GetValuesAround(x, y, 1)
	}

	newV := l.updateRule(box)

	l.SetValue(x, y, newV)
	return newV != box[1][1]
}

// SingleIteration of applying an update rule to a lattice. Returns true if any
// coordinate value in the lattice updated in the iteration.
func (l *Lattice[T]) SingleIteration() bool {
	isUpdated := false
	readL := l.Copy()
	for i := range l.GetLatticeCoordinates() {
		updated := l.UpdatePair(i, readL)
		if updated { isUpdated = true }
	}
	return isUpdated
}


// LatticeParams required to construct a Lattice with type uint (i.e. Lattice[uint]).
type LatticeParams struct {
	GridSize   uint
	AliveRatio float64
	Topology   string
}

func ConstructUintLattice(params LatticeParams, updateRule main.UpdateRuleFn) *Lattice[uint] {
	return &Lattice[uint]{
		grid:       ConstructUintGrid(params.GridSize, params.AliveRatio),
		topologyFn: GetTransformation(params.Topology),
		n:          params.GridSize,
		null:       uint(math.MaxUint),
		formatter: func(t uint) string {
			if t == 0 {
				return "-"
			} else {
				return "+"
			}
		},
		updateRule: updateRule,
	}
}

// ConstructUintGrid of binary values on a given square size `n`. Probability of a value of 1 is binaryProb.
func ConstructUintGrid(n uint, binaryProb float64) []uint {
	// Construct chan of 1/0s from decomposing uint64s.
	u := make(chan uint)
	go func(out chan uint, size uint) {
		for i := 0; i < int(size); i++ {
			if rand.Float64() < binaryProb { u <- 1 } else { u <- 0 }
		}
	}(u, n*n)

	rows := make([]uint, n*n)
	for i := 0; i < int(n*n); i++ {
		rows[i] = <- u
	}
	return rows
}
