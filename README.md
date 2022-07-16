# Go-L
Game of Life simulation 

I originally wrote a Game of Life model in C, [GOL](https://github.com/Jeadie/GOL/). But I wanted to convert it into Golang for a few reasons:
 - Learn Go
 - Consider more novel topoogies based on border rules (see #border-topologies).
 - Extend GOL simulation for generic update rules.
 - Extend GOL lattice for integer values (i.e. not just 1/0)
 - Run physical system simulations based on applications of cellular automaton:
   - Computational fluid dynamics
   - Population dynamics
   - Ising Models
   - Other interesting cellular automaton: Rule 90, Langton's ant
   - Boolean binary logic rules (i.e. consider the 1 cell + 4 neighbours as a 5 bit input - 1 bit output to a logic circuit). 
   - Discretise and extend to 2D my previous research in [rho signalling in cell-cell junctions during collective cell migration](https://github.com/Jeadie/UQ-Winter-Research-Project-2017).
   - 


## Usage
Current usage simply displays the simulation onto the terminal. Usage: 
```bash
> Go-L --help
Usage of Go-L:
  -aliveratio float
    	The fraction of squares that start as alive, assigned at random. Domain: [0.0, 1.0]. (default 0.8)
  -gridsize uint
    	Length of square grid to define game on. (default 20)
  -iterations uint
    	Max number of iterations to simulate game of life. If stable solution, will exit early. (default 100)
  -topology string
    	Specify the topology of the grid (as a fundamental topology from a parallelograms). Valid parameters: BORDERED, TORUS, KLEIN_BOTTLE, PROJECTIVE_PLANE, SPHERE. (default "BORDERED")
  -updatedelay uint
    	Additional period delay between updating rounds of the game, in milliseconds. Does not take into account processing time. (default 200)
```

### Border Topologies
In redefining how the border conditions work, we can simulate GOL as if it was played on a variety of manifolds. This is most clearly seen when looking at the fundamental polygons derived from a square (or parallelogram). When considering the neighbours of a cell on the border of the lattice, a fundamental polygon helps show where the neighbouring values should be.

![Sphere](https://en.wikipedia.org/wiki/File:SphereAsSquare.svg)
![Real Projective Plane](https://en.wikipedia.org/wiki/File:ProjectivePlaneAsSquare.svg)
![Klein Bottle](https://en.wikipedia.org/wiki/File:KleinBottleAsSquare.svg)
![Torus](https://en.wikipedia.org/wiki/File:TorusAsSquare.svg)

Consider a 5x5 lattice, coordinates indexed in [0,4] x [0,4]. For a standard coordinate, say (2,2), its neighbours are: {(1,2),(3,2),(2,1),(2,3)}. These are all within standard bounds. Now consider the point (4,2) with neighbours: {(3,2),(5,2),(4,1),(4,3)}. What value should we use for (5,2)? The topology dictates how lattice-border neighbours get selected. For a bordered topology, there is nothing outside of lattice, therefore index it the null value (or 0 in GOL rules). For a sphere, the (5,2) becomes (2,0) gives its equivalence relation. 

Current: 
- Bordered
- Sphere
- Torus

WIP: 
- KleinBottle
- ProjectivePlane
