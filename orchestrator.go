package main

import (
	"fmt"
)

type IntPair = [2]int

type BaseOrchestrator[T Node] interface {
	SingleIteration(l *Lattice[Node]) (bool, error)
}

type Orchestrator[T Node] struct {

}

func (o Orchestrator[Node]) SingleIteration(l *Lattice[Node], updateRule func([][]Node) Node) (bool, error) {
	fmt.Println("Running single iteration...")
	readG :=  make([]Node, len(l.grid))
	copied := copy(readG, l.grid)

	// Use for read only, to allow parallel write to original lattice.
	readL := &Lattice[Node]{
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

	// Update write Lattice using read Lattice.
	go func(write *Lattice[Node], read *Lattice[Node], ins chan IntPair) {
		for i := range ins {
			x, y := i[0], i[1]
			box := read.GetValuesAround(x, y, 1)
			write.SetValue(x, y, updateRule(box))
		}
	}(l, readL, x )


	return copied > 0, nil
}
