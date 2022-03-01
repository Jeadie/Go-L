package main

type Lattice struct {
}

func (l *Lattice) cleanup() {

}

type LatticeParams struct {
	gridSize   uint
	aliveRatio float64
	topology   string
}

func ConstructLattice(params LatticeParams) *Lattice {
	return &Lattice{}
}
