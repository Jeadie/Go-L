package main

import (
	"fmt"
	"math/rand"
)

// Node can be used in Lattice method signatures
type Node any
type Lattice[T Node] struct {
	grid [] T
	topology string
	n uint
}

func (l *Lattice[Node]) Cleanup() {
	fmt.Println("cleaning up Lattice...")
}

func (l *Lattice[Node]) GetValue(x int, y int) Node {
	nx, ny := TranslateVertex(x, y, 0,0, len(l.grid), l.topology)
	return l.grid[(nx * int(l.n)) + ny]
};

func (l *Lattice[Node]) SetValue(x int, y int, v Node)  {
	nx, ny := TranslateVertex(x, y, 0, 0, len(l.grid), l.topology)
	l.grid[(nx*int(l.n))+ny] = v
}

// GetValuesAround returns all values around a coordinate within an l-1 distance of w.
func (l *Lattice[Node]) GetValuesAround(x int, y int, w int) [][]Node {
	rows := make([][]Node, 2*w+1)

	for i := -w; i < w; i++ {
		rows[i+w] = make([]Node, 2*w + 1)
		for j := -w; j < w; j++ {
			rows[w+i][w+j] = l.GetValue(x+i, y+j)
		}
	}
	return rows
}

type LatticeParams struct {
	gridSize   uint
	aliveRatio float64
	topology   string
}

func ConstructUintLattice(params LatticeParams) *Lattice[uint] {
	return &Lattice[uint]{
		grid: ConstructUintGrid(params.gridSize, params.aliveRatio),
		topology: params.topology,
		n: params.gridSize,
	}
}

func ConstructUintGrid(n uint, binaryProb float64) []uint {
	// Construct chan of 1/0s from decomposing uint64s.
	u := make(chan uint)
	go func(out chan uint, size uint) {
		for i := 0; i < int(size); i++ {
			if rand.Float64() > binaryProb { u <- 1 } else { u <- 0 }
		}
	}(u, n*n)

	rows := make([]uint, n*n)
	for i := 0; i < int(n*n); i++ {
		rows[i] = <- u
	}
	return rows
}
