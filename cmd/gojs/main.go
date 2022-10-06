package main

import (
	"fmt"
	"github.com/Jeadie/Go-L/pkg/gol"
	//"github.com/teamortix/golang-wasm/wasm"
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

// This exports an add function.
// It takes in two 32-bit integer values
// And returns a 32-bit integer value.
// To make this function callable from JavaScript,
// we need to add the: "export add" comment above the function
//export add
func add(x int, y int) int {
	return x + y
}


func main() {
	//wasm.Expose("run", Run)
	//wasm.Ready()
	fmt.Println("Hello browser, I am from Golang world!")
}
