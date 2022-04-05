package main

import (
	"constraints"
)

type IntPair = [2]int

type BaseOrchestrator[T constraints.Ordered] interface {
	SingleIteration(l *Lattice[T]) (bool, error)
}

type Orchestrator[T constraints.Ordered] struct {

}

func (o Orchestrator[T]) SingleIteration(l *Lattice[T], updateRule func([][]T) T) (bool, error) {
	readG :=  make([]T, len(l.grid))
	copy(readG, l.grid)

	// Use for read only, to allow parallel write to original lattice.
	readL := &Lattice[T]{
		grid:     readG,
		topology: l.topology,
		n:        l.n,
	}

	// Get data from copied array
	x := make(chan IntPair)
	go func(out chan IntPair) {
		defer close(out)
		for i := 0; i < int(readL.n); i++ {
			for j := 0; j < int(readL.n); j++ {
				out <- [2]int{i, j}
			}
		}
	}(x)

	isUpdated := false
	// Update write Lattice using read Lattice.
	go func(write *Lattice[T], read *Lattice[T], ins chan IntPair, isUpdated *bool) {
		for i := range ins {
			x, y := i[0], i[1]
			box := read.GetValuesAround(x, y, 1)
			newV := updateRule(box)
			//fmt.Println(x, y, l.GetValue(x, y), newV, box)
			write.SetValue(x, y, newV)
			if newV != box[1][1] { *isUpdated = true }
		}
	}(l, readL, x, &isUpdated)

	l.Print()
	return isUpdated, nil
}
