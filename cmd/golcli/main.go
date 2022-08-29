package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Jeadie/Go-L/pkg/gol"
	"strings"
	"time"
)

type LatticeProcessor func(*gol.Lattice[uint]) error

func parseArguments() (*gol.InputParameters, error) {
	params := gol.InputParameters{}

	flag.UintVar(&params.Iterations, "iterations", 100, "Max number of iterations to simulate game of life. If stable solution, will exit early.")
	flag.UintVar(&params.GridSize, "gridsize", 20, "Length of square grid to define game on.")
	flag.Float64Var(&params.AliveRatio, "aliveratio", 0.8, "The fraction of squares that start as alive, assigned at random. Domain: [0.0, 1.0].")
	flag.UintVar(&params.UpdateDelay, "updatedelay", 200, "Additional period delay between updating rounds of the game, in milliseconds. Does not take into account processing time.")
	flag.StringVar(&params.Topology, "topology", gol.DefaultTopology, "Specify the topology of the grid (as a fundamental topology from a parallelograms). Valid parameters: BORDERED, TORUS, KLEIN_BOTTLE, PROJECTIVE_PLANE, SPHERE.")
	flag.UintVar(&params.UpdateFunctionNumber, "updaterule", gol.ConwaysGameOfLifeUpdateRuleNumber, "Specify the number associated with the update rule to use. Default to Conway's Game of Life.")
	flag.Parse()

	if gol.IsValidTopology(params.Topology) {
		return &params, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid topology specified %s. Topology must be one of %s", params.Topology, gol.ALLOWED_TOPOLOGIES))
}

func main() {
	// Parse cli args
	params, err := parseArguments()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup Lattice
	l := gol.ConstructUintLatticeFromInput(*params)

	processors := []LatticeProcessor{Print}

	fmt.Print("\033[H\033[2J")
	for i := uint(0); i < params.Iterations; i++ {
		if !l.SingleIteration() {break}
		runProcessors(l, processors)
	}
}

func Print(l *gol.Lattice[uint]) error {
	for i := 0; i <int(l.Size()); i++ {
		fmt.Printf("\033[%d;3H", i+2)
		line :=  make([]string, l.Size())
		for j := 0; j < int(l.Size()); j++ {
			line[j] = l.GetFormattedValueAt(j, i)
		}
		fmt.Println(strings.Join(line, " "))
	}
	time.Sleep(time.Millisecond * 400)
	return nil
}

// Run a set of read-only processes on a Lattice.
func runProcessors(l *gol.Lattice[uint], processors []LatticeProcessor) {
	for _, fn := range processors {
		if err := fn(l); err != nil {fmt.Println(err)}
	}
}