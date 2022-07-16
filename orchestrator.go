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

func (o Orchestrator[T]) UpdateLattice(l *Lattice[T], readL *Lattice[T], pairs chan IntPair) bool {
	isUpdated := false
	for i := range pairs {
		updated := o.UpdatePair(i, readL, l)
		if updated { isUpdated = true }
	}
	return isUpdated
}

func (o Orchestrator[T]) UpdatePair(i IntPair, readL *Lattice[T], writeL *Lattice[T]) bool {
	x, y := i[0], i[1]
	box := readL.GetValuesAround(x, y, 1)
	newV := writeL.updateRule(box)

	writeL.SetValue(x, y, newV)
	return newV != box[1][1]
}

func (o Orchestrator[T]) SingleIteration(l *Lattice[T]) (bool, error) {
	// Use for read only, to allow parallel write to original lattice.
	readL := l.Duplicate()
	x := readL.GetLatticeCoordinates()
	return o.UpdateLattice(l, readL, x), nil
}
