package main

import "fmt"

type BaseOrchestrator[T Node] interface {
	SingleIteration(l *Lattice[Node]) (bool, error)
}

type Orchestrator[T Node] struct {
}

func (o Orchestrator[Node]) SingleIteration(l *Lattice[Node]) (bool, error) {
	fmt.Println("Running single iteration...")
	return true, nil
}
