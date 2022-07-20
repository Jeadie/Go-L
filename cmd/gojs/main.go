package main

import (
	"github.com/Jeadie/Go-L/pkg/gol"
	"github.com/teamortix/golang-wasm/wasm"
)

// Run a cellular automaton simulation. Return type is mapping of iteration index to grid values.
func Run(params gol.InputParameters) map[int][][]uint {

	l := gol.ConstructUintLatticeFromInput(params)
	output := make(map[int][][]uint, params.Iterations)

	for i := uint(0); i < params.Iterations; i++ {
		output[int(i)] = l.MakeGrid()
		if !l.SingleIteration() {break}
	}
	return output
}

func main() {
	wasm.Expose("run", Run)
	wasm.Ready()
}
