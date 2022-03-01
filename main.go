package main

import (
	"flag"
	"fmt"
)

type InputParameters struct {
	iterations uint
	gridSize uint
	aliveRatio float64
	updateDelay uint
}

func parseArguments() *InputParameters {
	params := InputParameters{}

	flag.UintVar(&params.iterations, "iterations", 100, "Max number of iterations to simulate game of life. If stable solution, will exit early.")
	flag.UintVar(&params.gridSize, "gridsize", 20, "Length of square grid to define game on.")
	flag.Float64Var(&params.aliveRatio, "aliveratio", 0.8, "The fraction of squares that start as alive, assigned at random. Domain: [0.0, 1.0]")
	flag.UintVar(&params.updateDelay, "updatedelay", 200,"Additional period delay between updating rounds of the game, in milliseconds. Does not take into account processing time.")
	flag.Parse()

	return &params
}

func main() {
	params := parseArguments()
	fmt.Println("Hello World", params)
}
