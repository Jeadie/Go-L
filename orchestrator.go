package main

type BaseOrchestrator[T Node] interface {
	SingleIteration(l *Lattice[Node]) (bool, error)
}

type Orchestrator[T Node] struct {
}

func (o Orchestrator[Node]) SingleIteration(l *Lattice[Node]) (bool, error) {
	return true, nil
}
