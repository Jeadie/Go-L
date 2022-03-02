package main


// Node can be used in Lattice method signatures
type Node any
type Lattice[T Node] struct {
	grid [][] T
}

func (l *Lattice[Node]) Cleanup() {

}

func (l *Lattice[Node]) GetValue(x int, y int) Node {
	return l.grid[x][y]
};

func (l *Lattice[Node]) GetValuesAround(x int, y int, w int) [][]Node {
	return l.grid[x-w:x+w][y-w:y+w]
}

type LatticeParams struct {
	gridSize   uint
	aliveRatio float64
	topology   string
}

func ConstructLattice[S Node](params LatticeParams) *Lattice[S] {
	return &Lattice[S]{
		grid: [][]S{},
	}
}
