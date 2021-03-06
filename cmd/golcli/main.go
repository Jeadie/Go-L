package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Jeadie/Go-L/pkg/gol"
	"strings"
	"time"
)

type InputParameters struct {
	iterations  uint
	gridSize    uint
	aliveRatio  float64
	updateDelay uint
	topology    string
	updateFunctionNumber uint
}

const ConwaysGameOfLifeUpdateRuleNumber = 1994975360

type LatticeProcessor func(*gol.Lattice[uint]) error

func parseArguments() (*InputParameters, error) {
	params := InputParameters{}

	flag.UintVar(&params.iterations, "iterations", 100, "Max number of iterations to simulate game of life. If stable solution, will exit early.")
	flag.UintVar(&params.gridSize, "gridsize", 20, "Length of square grid to define game on.")
	flag.Float64Var(&params.aliveRatio, "aliveratio", 0.8, "The fraction of squares that start as alive, assigned at random. Domain: [0.0, 1.0].")
	flag.UintVar(&params.updateDelay, "updatedelay", 200, "Additional period delay between updating rounds of the game, in milliseconds. Does not take into account processing time.")
	flag.StringVar(&params.topology, "topology", gol.DefaultTopology, "Specify the topology of the grid (as a fundamental topology from a parallelograms). Valid parameters: BORDERED, TORUS, KLEIN_BOTTLE, PROJECTIVE_PLANE, SPHERE.")
	flag.UintVar(&params.updateFunctionNumber, "updaterule", ConwaysGameOfLifeUpdateRuleNumber, "Specify the number associated with the update rule to use. Default to Conway's Game of Life.")
	flag.Parse()

	if gol.IsValidTopology(params.topology) {
		return &params, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid topology specified %s. Topology must be one of %s", params.topology, gol.ALLOWED_TOPOLOGIES))
}

func main() {
	// Parse cli args
	params, err := parseArguments()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Update rule
	var updateRule gol.UpdateRuleFn
	if params.updateFunctionNumber == ConwaysGameOfLifeUpdateRuleNumber {
		updateRule = gol.CalculateGOLValue
	} else {
		updateRule = gol.CreateUpdateRule(params.updateFunctionNumber)
	}

	// Setup Lattice
	l := gol.ConstructUintLattice(
		gol.LatticeParams{
			GridSize:   params.gridSize,
			AliveRatio: params.aliveRatio,
			Topology:   params.topology,
		},
		updateRule,
	)

	processors := []LatticeProcessor{Print}

	fmt.Print("\033[H\033[2J")
	for i := uint(0); i < params.iterations; i++ {
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