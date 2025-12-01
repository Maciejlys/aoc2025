package utils

// Point represents a 2D coordinate
type Point struct {
	X, Y int
}

// Add returns a new point that is the sum of p and other
func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

// Sub returns a new point that is the difference of p and other
func (p Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

// ManhattanDistance returns the Manhattan distance to another point
func (p Point) ManhattanDistance(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// Neighbors4 returns the 4 orthogonal neighbors (up, down, left, right)
func (p Point) Neighbors4() []Point {
	return []Point{
		{X: p.X, Y: p.Y - 1}, // up
		{X: p.X, Y: p.Y + 1}, // down
		{X: p.X - 1, Y: p.Y}, // left
		{X: p.X + 1, Y: p.Y}, // right
	}
}

// Neighbors8 returns all 8 neighbors including diagonals
func (p Point) Neighbors8() []Point {
	return []Point{
		{X: p.X - 1, Y: p.Y - 1}, // top-left
		{X: p.X, Y: p.Y - 1},     // top
		{X: p.X + 1, Y: p.Y - 1}, // top-right
		{X: p.X - 1, Y: p.Y},     // left
		{X: p.X + 1, Y: p.Y},     // right
		{X: p.X - 1, Y: p.Y + 1}, // bottom-left
		{X: p.X, Y: p.Y + 1},     // bottom
		{X: p.X + 1, Y: p.Y + 1}, // bottom-right
	}
}

// Direction vectors for common movement
var (
	Up    = Point{X: 0, Y: -1}
	Down  = Point{X: 0, Y: 1}
	Left  = Point{X: -1, Y: 0}
	Right = Point{X: 1, Y: 0}
)
