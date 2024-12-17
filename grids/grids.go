package grids

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

type Point struct {X, Y int}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p *Point) Direction(other *Point) int {
	if p.X != other.X && p.Y != other.Y {
		panic("Points do not align!")
	}

	if other.X > p.X {
		return EAST
	}
	if other.X < p.X {
		return WEST
	}
	if other.Y > p.Y {
		return SOUTH
	}
	if other.Y < p.Y {
		return NORTH
	}

	panic("Unable to calculate direction")
}

type Grid[T comparable] struct {
	grid map[Point]T
	maxX, maxY int
}

func NewGrid[T comparable]() (*Grid[T]) {
	return &Grid[T]{grid: make(map[Point]T)}
}

func (g *Grid[T]) Height() int {
	return g.maxY
}

func (g *Grid[T]) Width() int {
	return g.maxX
}

func (g *Grid[T]) AddPoint(x int, y int, item T) {
	g.grid[Point{x, y}] = item
	if x > g.maxX {
		g.maxX = x
	}
	if y > g.maxY {
		g.maxY = y
	}
}

func (g *Grid[T]) AddP(point Point, item T) {
	g.grid[point] = item
	if point.X > g.maxX {
		g.maxX = point.X
	}
	if point.Y > g.maxY {
		g.maxY = point.Y
	}
}

func (g *Grid[T]) Get(point Point) (T, bool) {
	item, ok := g.grid[point]
	return item, ok
}

func (g *Grid[T]) Set(point Point, value T) {
	g.grid[point] = value
}

// Returns each point in the grid by traversing it left to right, top to bottom
func (g *Grid[T]) Iter() (chan Point) {
	c := make(chan Point)

	go func() {
		x := 0
		y := 0
		for {
			if x > g.maxX {
				x = 0
				y++
				continue
			}

			if y > g.maxY {
				break
			}

			c <- Point{x, y}
			x++
		}
		close(c)
	}()

	return c
}

// Returns the points in the 4 cardinal directions from `point`.
// If one does not exist in `grid` it is not returned.
func (g *Grid[T]) Directions(point Point) (dirs []Point) {
	dirs = make([]Point, 0)
	points := []Point{
		Point{point.X, point.Y-1},
		Point{point.X+1, point.Y},
		Point{point.X, point.Y+1},
		Point{point.X-1, point.Y},
	}

	for _, p := range points {
		if _, ok := g.grid[p]; ok {
			dirs = append(dirs, p)
		}
	}

	return
}

func (g *Grid[any]) Size() int {
	return len(g.grid)
}

func (g *Grid[T]) Lines() [][]T {
	x := 0
	y := 0
	lines := make([][]T, g.maxY + 1)
	line := make([]T, g.maxX + 1)
	for {
		if x > g.maxX {
			x = 0
			lines[y] = line
			line = make([]T, g.maxX + 1)
			y++
			continue
		}

		if y > g.maxY {
			break
		}

		line[x] = g.grid[Point{x, y}]
		x++
	}

	return lines
}
