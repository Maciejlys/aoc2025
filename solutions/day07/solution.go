package day07

import (
	"github.com/Maciejlys/aoc2025/utils"
)

func Part1(input string) int {
	parsed := utils.Parse(input, "")
	grid := make(map[utils.Point]string)
	beams := utils.NewQueue[utils.Point]()
	sum := 0
	var start utils.Point

	for y, line := range parsed {
		for x, r := range line {
			grid[utils.Point{X: x, Y: y}] = r
			if r == "S" {
				start = utils.Point{X: x, Y: y}
				beams.Enqueue(start)
			}
		}
	}

	visited := make(map[utils.Point]bool)

	for !beams.IsEmpty() {
		current, _ := beams.Dequeue()

		if visited[current] {
			continue
		}

		visited[current] = true

		if grid[current] == "" {
			continue
		}

		if grid[current] == "^" {
			sum++
			for _, p := range current.Neighbors2LR() {
				beams.Enqueue(p)
			}
			continue
		}

		moved := current.Add(utils.Down)
		beams.Enqueue(moved)
	}

	return sum
}

type State struct {
	position utils.Point
	char     string
	beams    int
}

func Part2(input string) int {
	parsed := utils.Parse(input, "")
	grid := make(map[utils.Point]State)

	for y, line := range parsed {
		for x, r := range line {
			p := utils.Point{X: x, Y: y}
			beams := 0
			if r == "S" {
				beams++
			}
			grid[p] = State{position: p, beams: beams, char: r}
		}
	}

	sum := 0

	for {
		hasBeams := false

		for y := range len(parsed) {
			for x := range len(parsed[0]) {
				p := utils.Point{X: x, Y: y}
				state := grid[p]

				if state.beams == 0 {
					continue
				}

				hasBeams = true
				beamsToPropagate := state.beams

				state.beams = 0
				grid[p] = state

				if state.char == "^" {
					sum += beamsToPropagate

					for _, neighbor := range p.Neighbors2LR() {
						if entry, ok := grid[neighbor]; ok {
							entry.beams += beamsToPropagate
							grid[neighbor] = entry
						}
					}
				} else {
					down := p.Add(utils.Down)
					if entry, ok := grid[down]; ok {
						entry.beams += beamsToPropagate
						grid[down] = entry
					}
				}
			}
		}

		if !hasBeams {
			break
		}
	}

	return sum + 1
}
