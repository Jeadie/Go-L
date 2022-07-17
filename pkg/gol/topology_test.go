package gol

import (
	"fmt"
	"testing"
)

type TopologyTestCase struct {
	x, y, dx, dy, n      int
	expectedX, expectedY int
}

type TransformTestSet struct {
	tests    []TopologyTestCase
	topology string
}

func RunCase(test TopologyTestCase, fn TopologyTransformation) error {
	nx, ny := fn(test.x, test.y, test.dx, test.dy, test.n)
	if nx != test.expectedX || ny != test.expectedY {
		return fmt.Errorf(
			"incorrect output for (%d+%d, %d+%d) on size %d. Expected (%d, %d), output (%d, %d)",
			test.x, test.dx, test.y, test.dy, test.n, test.expectedX, test.expectedY, nx, ny)
	}
	return nil
}

func TestTopologyTransformation(t *testing.T) {
	tests := []TransformTestSet{{
		topology: Bordered,
		tests: []TopologyTestCase{
			{x: 0, y: 0, dx: -1, dy: 0, n: 5, expectedX: -1, expectedY: -1},
			{x: 0, y: 0, dx: 0, dy: -1, n: 5, expectedX: -1, expectedY: -1},
			{x: 0, y: 0, dx: -1, dy: -1, n: 5, expectedX: -1, expectedY: -1},
			{x: 0, y: 0, dx: 0, dy: 1, n: 5, expectedX: 0, expectedY: 1},
		},
	}, {
		topology: Torus,
		tests: []TopologyTestCase{
			{x: 0, y: 0, dx: -1, dy: 0, n: 5, expectedX: 4, expectedY: 0},
			{x: 0, y: 0, dx: 0, dy: -1, n: 5, expectedX: 0, expectedY: 4},
			{x: 0, y: 0, dx: -1, dy: -1, n: 5, expectedX: 4, expectedY: 4},
			{x: 0, y: 0, dx: 0, dy: 1, n: 5, expectedX: 0, expectedY: 1},
		},
	}, {
		topology: KleinBottle,
		tests: []TopologyTestCase{
			// Boundary Conditions
			{x: 5, y: 1, dx: 0, dy: 0, n: 5, expectedX: 0, expectedY: 1},
			{x: 3, y: 5, dx: 0, dy: 0, n: 5, expectedX: 2, expectedY: 0},

			// Corner conditions
			{x: 5, y: 0, dx: 0, dy: 0, n: 5, expectedX: 0, expectedY: 0},
			{x: 0, y: 5, dx: 0, dy: 0, n: 5, expectedX: 0, expectedY: 0},
			{x: 5, y: 5, dx: 0, dy: 0, n: 5, expectedX: 0, expectedY: 0},

			// Out of bound conditions
			{x: 0, y: 0, dx: -1, dy: 0, n: 5, expectedX: 4, expectedY: 0},
			{x: 0, y: 0, dx: 0, dy: -1, n: 5, expectedX: 0, expectedY: 4},
			{x: 0, y: 0, dx: -1, dy: -1, n: 5, expectedX: 4, expectedY: 4},
			{x: 0, y: 0, dx: 0, dy: 1, n: 5, expectedX: 0, expectedY: 1},
		},
	}, {
		topology: ProjectivePlane,
		tests: []TopologyTestCase{

			// Boundary conditions
			{x: 3, y: 5, dx: 0, dy: 0, n: 5, expectedX: 2, expectedY: 0},
			{x: 5, y: 3, dx: 0, dy: 0, n: 5, expectedX: 0, expectedY: 2},

			// Corner conditions
			{x: 5, y: 5, dx: 0, dy: 0, n: 5, expectedX: 0, expectedY: 0},

			// These are unmappable, default to (-1, -1)
			{x: 5, y: 0, dx: 0, dy: 0, n: 5, expectedX: -1, expectedY: -1},
			{x: 0, y: 5, dx: 0, dy: 0, n: 5, expectedX: -1, expectedY: -1},

			// Out of bound conditions
			{x: 0, y: 0, dx: -1, dy: 0, n: 5, expectedX: -1, expectedY: -1},
			{x: 0, y: 0, dx: 0, dy: -1, n: 5, expectedX: -1, expectedY: -1},
			{x: 0, y: 0, dx: -1, dy: -1, n: 5, expectedX: -1, expectedY: -1},
			{x: 0, y: 0, dx: 0, dy: 1, n: 5, expectedX: 0, expectedY: 1},
		},
	}, {
		topology: Sphere,
		tests: []TopologyTestCase{
			{x: 0, y: 0, dx: -1, dy: 0, n: 5, expectedX: 1, expectedY: 0},
			{x: 0, y: 0, dx: 0, dy: -1, n: 5, expectedX: 0, expectedY: 1},
			{x: 0, y: 0, dx: -1, dy: -1, n: 5, expectedX: 4, expectedY: 4},

			{x: 3, y: 4, dx: 0, dy: 1, n: 5, expectedX: 0, expectedY: 3},
			{x: 4, y: 3, dx: 1, dy: 0, n: 5, expectedX: 3, expectedY: 0},
			{x: 4, y: 4, dx: 1, dy: 1, n: 5, expectedX: 0, expectedY: 0},

			{x: 0, y: 0, dx: 0, dy: 1, n: 5, expectedX: 0, expectedY: 1},
		},
	},
	}

	for _, testCase := range tests {
		t.Run(testCase.topology, func(t *testing.T) {
			fn := GetTransformation(testCase.topology)
			for _, test := range testCase.tests {
				err := RunCase(test, fn)
				if err != nil {
					t.Errorf(err.Error())
				}
			}
		})
	}
}
