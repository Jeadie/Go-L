package main

type BaseOrchestrator interface {
	SingleIteration(l *Lattice) (bool, error)
}

type Orchestrator struct {
}

func (o Orchestrator) SingleIteration(l *Lattice) (bool, error) {
	return true, nil
}
