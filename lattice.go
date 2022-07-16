package main

import (
	"constraints"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// T can be used in Lattice method signatures
type Node constraints.Ordered

type Lattice[T Node] struct {
	grid [] T
	topology string
	n uint
	null T
	formatter func(x T) string
	updateRule func([][]T) T
}

func (l *Lattice[T]) Cleanup() {}

func (l *Lattice[T]) GetValue(x int, y int) T {
	nx, ny := TranslateVertex(x, y, 0,0, int(l.n), l.topology)
	if nx == -1 || ny == -1 { return l.null }
	return l.grid[(nx * int(l.n)) + ny]
}

func (l *Lattice[T]) SetValue(x int, y int, v T)  {
	nx, ny := TranslateVertex(x, y, 0, 0, len(l.grid), l.topology)
	l.grid[(nx*int(l.n))+ny] = v
}

func (l *Lattice[T]) Print() {
	for i := 0; i <int(l.n); i++ {
		fmt.Printf("\033[%d;3H", i+2)
		line :=  make([]string, l.n)
		for j := 0; j < int(l.n); j++ {
			line[j] = l.formatter(l.GetValue(j, i))
		}
		fmt.Println(strings.Join(line, " "))
	}
	time.Sleep(time.Millisecond * 400)
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

func (l *Lattice[T]) Duplicate() *Lattice[T] {
	readG :=  make([]T, len(l.grid))
	copy(readG, l.grid)
	return &Lattice[T]{
		grid:     readG,
		topology: l.topology,
		n:        l.n,
		updateRule: l.updateRule,
	}
}

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


func (l *Lattice[T]) UpdateLattice(readOnlyL *Lattice[T], pairs chan IntPair) bool {
	isUpdated := false
	for i := range pairs {
		updated := l.UpdatePair(i, readOnlyL)
		if updated { isUpdated = true }
	}
	return isUpdated
}

func (l *Lattice[T]) UpdatePair(i IntPair, readL *Lattice[T]) bool {
	x, y := i[0], i[1]
	box := readL.GetValuesAround(x, y, 1)
	newV := l.updateRule(box)

	l.SetValue(x, y, newV)
	return newV != box[1][1]
}

func (l *Lattice[T]) SingleIteration() bool {
	return l.UpdateLattice(
		l.Duplicate(),
		l.GetLatticeCoordinates(),
	)
}


// LatticeParams required to construct a Lattice with type uint (i.e. Lattice[uint]).
type LatticeParams struct {
	gridSize   uint
	aliveRatio float64
	topology   string
}

func ConstructUintLattice(params LatticeParams, updateRule func([][]uint) uint) *Lattice[uint] {
	return &Lattice[uint]{
		grid: ConstructUintGrid(params.gridSize, params.aliveRatio),
		topology: params.topology,
		n: params.gridSize,
		null: uint(math.MaxUint),
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
