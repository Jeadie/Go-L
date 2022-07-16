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
| Sphere | Real Projective Plane | Klein Bottle | Torus | 
| -- | -- | -- | -- | 
| ![Sphere](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c1/SphereAsSquare.svg/240px-SphereAsSquare.svg.png)| ![Real Projective Plane](https://upload.wikimedia.org/wikipedia/commons/thumb/9/9b/ProjectivePlaneAsSquare.svg/240px-ProjectivePlaneAsSquare.svg.png)| ![Klein Bottle](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/KleinBottleAsSquare.svg/240px-KleinBottleAsSquare.svg.png)| ![Torus](https://upload.wikimedia.org/wikipedia/commons/thumb/f/f2/TorusAsSquare.svg/240px-TorusAsSquare.svg.png) |


Consider a square lattice of size 5, coordinates indexed in $[0,4] \times [0,4]$. For a standard coordinate, say $(2,2)$, its neighbours are: ${(1,2),(3,2),(2,1),(2,3)}$. These are all within standard bounds. Now consider the point $(4,2)$ with neighbours: ${(3,2),(5,2),(4,1),(4,3)}$. What value should we use for $(5,2)$? The topology dictates how lattice-border neighbours get selected. For a bordered topology, there is nothing outside of lattice, therefore index it the null value (or 0 in GOL rules). For a sphere, the $(5,2)$ becomes $(2,0)$ given its equivalence relation (on an $n$ square lattice): 

$$ 
(x, 0) \backsim (n, n-x), \quad x\in [0, n] \\
(x, n) \backsim (0, n-x), \quad x\in [0, n]
$$
Or for a torus
$$ 
(x, 0) \backsim (x, n), \quad x\in [0, n] \\
(0, y) \backsim (n, y), \quad y\in [0, n]
$$
or a real projective plane
$$ 
(0, y) \backsim (n, n-y), \quad y\in [0, n] \\
(x, 0) \backsim (n-x, n), \quad x\in [0, n]
$$
and lastly a klein bottle
$$ 
(0, y) \backsim (n, y), \quad y\in [0, n] \\
(x, 0) \backsim (n-x, n), \quad x\in [0, n]
$$

For update rules that consider 2nd degree neighbours (i.e. $(6,2)$), the mapping gets a bit more complicated. 