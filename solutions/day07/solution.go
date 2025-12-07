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
	H := len(parsed)
	W := len(parsed[0])

	grid := make(map[utils.Point]State)

	for y, line := range parsed {
		for x, r := range line {
			p := utils.Point{X: x, Y: y}
			grid[p] = State{position: p, beams: 0, char: r}
			if r == "S" {
				grid[p] = State{position: p, beams: 1, char: r}
			}
		}
	}

	sum := 0

	for {
		nextGrid := make(map[utils.Point]State)
		for p, state := range grid {
			nextGrid[p] = State{position: p, char: state.char, beams: 0}
		}

		hasBeams := false

		for y := range H {
			for x := range W {
				p := utils.Point{X: x, Y: y}
				state := grid[p]

				if state.beams == 0 {
					continue
				}

				hasBeams = true
				beamsToPropagate := state.beams

				if state.char == "^" {
					sum += beamsToPropagate

					left := p.Add(utils.Left)
					right := p.Add(utils.Right)

					if entry, ok := nextGrid[left]; ok {
						entry.beams += beamsToPropagate
						nextGrid[left] = entry
					}

					if entry, ok := nextGrid[right]; ok {
						entry.beams += beamsToPropagate
						nextGrid[right] = entry
					}

					continue
				}

				down := p.Add(utils.Down)
				if entry, ok := nextGrid[down]; ok {
					entry.beams += beamsToPropagate
					nextGrid[down] = entry
				}
			}
		}

		if !hasBeams {
			break
		}

		grid = nextGrid
	}

	return sum + 1
}
