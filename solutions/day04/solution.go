package day04

import (
	"github.com/Maciejlys/aoc2025/utils"
)

func getGrid(input string) map[utils.Point]rune {
	grid := make(map[utils.Point]rune, 0)
	lines := utils.Lines(input)

	for y, line := range lines {
		for x, char := range line {
			grid[utils.Point{X: x, Y: y}] = char
		}
	}

	return grid
}

func Part1(input string) int {
	grid := getGrid(input)
	sum := 0

	for point := range grid {
		neighbors := point.Neighbors8()
		counter := 0

		if grid[point] != '@' {
			continue
		}

		for _, neighbor := range neighbors {
			if grid[neighbor] == '@' {
				counter++
			}
		}

		if counter < 4 {
			sum++
		}
	}

	return sum
}

func Part2(input string) int {
	grid := getGrid(input)
	stack := utils.NewStack[utils.Point]()
	sum := 0

	for !stack.IsEmpty() || sum == 0 {
		sum += stack.Size()

		for !stack.IsEmpty() {
			point, _ := stack.Pop()
			grid[point] = '.'
		}

		for point := range grid {
			neighbors := point.Neighbors8()
			counter := 0

			if grid[point] != '@' {
				continue
			}

			for _, neighbor := range neighbors {
				if grid[neighbor] == '@' {
					counter++
				}
			}

			if counter < 4 {
				stack.Push(point)
			}
		}
	}

	return sum
}
