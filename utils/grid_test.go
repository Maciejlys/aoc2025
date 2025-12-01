package utils

import "testing"

func TestPointAdd(t *testing.T) {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 3, Y: 4}
	result := p1.Add(p2)

	if result.X != 4 || result.Y != 6 {
		t.Errorf("Point.Add() = %v, want {4, 6}", result)
	}
}

func TestPointSub(t *testing.T) {
	p1 := Point{X: 5, Y: 7}
	p2 := Point{X: 2, Y: 3}
	result := p1.Sub(p2)

	if result.X != 3 || result.Y != 4 {
		t.Errorf("Point.Sub() = %v, want {3, 4}", result)
	}
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		p1, p2   Point
		expected int
	}{
		{Point{0, 0}, Point{3, 4}, 7},
		{Point{1, 1}, Point{1, 1}, 0},
		{Point{-2, 3}, Point{2, -1}, 8},
	}

	for _, tt := range tests {
		result := tt.p1.ManhattanDistance(tt.p2)
		if result != tt.expected {
			t.Errorf("%v.ManhattanDistance(%v) = %d, want %d", tt.p1, tt.p2, result, tt.expected)
		}
	}
}

func TestNeighbors4(t *testing.T) {
	p := Point{X: 5, Y: 5}
	neighbors := p.Neighbors4()

	if len(neighbors) != 4 {
		t.Errorf("Neighbors4() returned %d neighbors, want 4", len(neighbors))
	}

	expected := []Point{
		{5, 4}, // up
		{5, 6}, // down
		{4, 5}, // left
		{6, 5}, // right
	}

	for i, n := range neighbors {
		if n != expected[i] {
			t.Errorf("Neighbor %d = %v, want %v", i, n, expected[i])
		}
	}
}

func TestNeighbors8(t *testing.T) {
	p := Point{X: 5, Y: 5}
	neighbors := p.Neighbors8()

	if len(neighbors) != 8 {
		t.Errorf("Neighbors8() returned %d neighbors, want 8", len(neighbors))
	}
}
