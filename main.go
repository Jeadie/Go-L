package main

import (
	"errors"
	"flag"
	"fmt"
)

const (
	Bordered        string = "BORDERED"
	Torus                  = "TORUS"
	Sphere				   = "SPHERE"
	KleinBottle            = "KLEIN_BOTTLE"
	ProjectivePlane        = "PROJECTIVE_PLANE"
)

var ALLOWED_TOPOLOGIES = []string{Bordered, Torus, KleinBottle, ProjectivePlane, Sphere}

func isValidTopology(x string) bool {
	for _, t := range ALLOWED_TOPOLOGIES {
		if t == x {
			return true
		}
	}
	return false
}

type InputParameters struct {
	iterations  uint
	gridSize    uint
	aliveRatio  float64
	updateDelay uint
	topology    string
}

func parseArguments() (*InputParameters, error) {
	params := InputParameters{}

	flag.UintVar(&params.iterations, "iterations", 100, "Max number of iterations to simulate game of life. If stable solution, will exit early.")
	flag.UintVar(&params.gridSize, "gridsize", 20, "Length of square grid to define game on.")
	flag.Float64Var(&params.aliveRatio, "aliveratio", 0.8, "The fraction of squares that start as alive, assigned at random. Domain: [0.0, 1.0].")
	flag.UintVar(&params.updateDelay, "updatedelay", 200, "Additional period delay between updating rounds of the game, in milliseconds. Does not take into account processing time.")
	flag.StringVar(&params.topology, "topology", DefaultTopology, "Specify the topology of the grid (as a fundamental topology from a parallelograms). Valid parameters: BORDERED, TORUS, KLEIN_BOTTLE, PROJECTIVE_PLANE, SPHERE.")
	flag.Parse()

	if isValidTopology(params.topology) {
		return &params, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid topology specified %s. Topology must be one of %s", params.topology, ALLOWED_TOPOLOGIES))
}

func main() {
	params, err := parseArguments()
	if err != nil {
		fmt.Println(err)
		return
	}
	lattice := ConstructUintLattice(LatticeParams{
		gridSize:   params.gridSize,
		aliveRatio: params.aliveRatio,
		topology:   params.topology,
	})
	latticeHandlers := getLatticeHandlers()
	o := Orchestrator[uint]{}
	defer lattice.Cleanup()

	fmt.Print("\033[H\033[2J")
	for i := uint(0); i < params.iterations; i++ {
		isChanged, err := o.SingleIteration(lattice, CalculateGOLValue)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !isChanged {
			break
		}
		for _, fn := range latticeHandlers {
			if err := fn(lattice); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func emptyHandler(l *Lattice[uint]) error {
	return nil
}

func getLatticeHandlers() []func(*Lattice[uint]) error {
	return []func(*Lattice[uint]) error{emptyHandler}
}
