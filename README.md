mars-rover-kata-golang
======================

[Mars Rover Kata](http://craftsmanship.sv.cmu.edu/katas/mars-rover-kata) with [Go](http://golang.org/)  

 * Develop an api that moves a rover around on a grid.
 * You are given the initial starting point (x,y) of a rover and the direction (N,S,E,W) it is facing.
 * The rover receives a character array of commands.
 * Implement commands that move the rover forward/backward (f,b).
 * Implement commands that turn the rover left/right (l,r).
 * Implement wrapping from one edge of the grid to another. (planets are spheres after all)
 * Implement obstacle detection before each move to a new square. If a given sequence of commands encounters an obstacle, the rover moves up to the last possible point and reports the obstacle.
 * Example: The rover is on a 100x100 grid at location (0, 0) and facing NORTH. The rover is given the commands "ffrff" and should end up at (2, 2)
 * Tips: use multiple classes and TDD

## Installing Go
see http://golang.org/doc/install or

```
brew install go
```

## Usage
```
git clone https://github.com/mattes/mars-rover-kata-golang.git
cd mars-rover-kata-golang
export GOPATH=/absolute/path/to/mars-rover-kata-golang # replace with your absolute path
go test rover # run tests
```

## Coordinates

```
(x,y)

                         (0,1)
                        NORTH: 0°

                           | y
                           |
                           | (0,0)
(-1,0) WEST: 270°   -----------------  EAST: 90° (1,0)
                           |       x  
                           |
                           |

                        SOUTH: 180°
                         (0,-1)
       

```
