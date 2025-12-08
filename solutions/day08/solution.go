package day08

import (
	"fmt"
	"sort"

	"github.com/Maciejlys/aoc2025/utils"
)

type Vec3 struct {
	X, Y, Z int
}

type Edge struct {
	i, j, d int
}

func (v Vec3) Dist(other Vec3) int {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z

	return dx*dx + dy*dy + dz*dz
}

func find(parent []int, x int) int {
	for parent[x] != x {
		parent[x] = parent[parent[x]]
		x = parent[x]
	}
	return x
}

func union(parent, size []int, a, b int) bool {
	ra, rb := find(parent, a), find(parent, b)
	if ra == rb {
		return false
	}
	if size[ra] < size[rb] {
		ra, rb = rb, ra
	}
	parent[rb] = ra
	size[ra] += size[rb]
	return true
}

func parseInput(input string) []Vec3 {
	lines := utils.Lines(input)
	var points []Vec3

	for _, l := range lines {
		var x, y, z int
		fmt.Sscanf(l, "%d,%d,%d", &x, &y, &z)
		points = append(points, Vec3{x, y, z})
	}

	return points
}

func createEdges(points []Vec3, n int) []Edge {
	var edges []Edge

	for i := range n {
		for j := i + 1; j < n; j++ {
			edges = append(edges, Edge{i, j, points[i].Dist(points[j])})
		}
	}

	sort.Slice(edges, func(a, b int) bool { return edges[a].d < edges[b].d })
	return edges
}

func Part1(input string) int {
	points := parseInput(input)
	n := len(points)
	edges := createEdges(points, n)

	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	for k := 0; k < 1000 && k < len(edges); k++ {
		union(parent, size, edges[k].i, edges[k].j)
	}

	var sizes []int
	for i := range n {
		if find(parent, i) == i {
			sizes = append(sizes, size[i])
		}
	}
	sort.Slice(sizes, func(a, b int) bool { return sizes[a] > sizes[b] })

	return sizes[0] * sizes[1] * sizes[2]
}

func Part2(input string) int {
	points := parseInput(input)
	n := len(points)
	edges := createEdges(points, n)

	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	components := n
	var last Edge

	for _, e := range edges {
		if union(parent, size, e.i, e.j) {
			last = e
			components--
			if components == 1 {
				break
			}
		}
	}

	return points[last.i].X * points[last.j].X
}
