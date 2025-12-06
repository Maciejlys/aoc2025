package day06

import (
	"strings"

	"github.com/Maciejlys/aoc2025/utils"
)

func parseInput(input string) [][]string {
	lines := utils.Lines(input)
	parsed := make([][]string, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		parsed = append(parsed, fields)
	}

	return parsed
}

func getColumns(parsed [][]string) map[int][]string {
	maths := make(map[int][]string)

	for _, line := range parsed {
		for i, char := range line {
			maths[i] = append(maths[i], string(char))
		}
	}

	return maths
}

func Part1(input string) int {
	parsed := parseInput(input)
	columns := getColumns(parsed)
	sum := 0

	for _, col := range columns {
		operation := col[len(col)-1]
		nums := col[:len(col)-1]
		converted := utils.StringsToInts(nums)

		if operation == "+" {
			sum += utils.Sum(converted)
		} else {
			sum += utils.Product(converted)
		}
	}

	return sum
}

func extractOperations(line []string) []string {
	operations := make([]string, 0)

	for _, op := range line {
		if op == "+" || op == "*" {
			operations = append(operations, op)
		}
	}

	return operations
}

func reorderNumbers(grid [][]string) map[int][]string {
	colIndex := 0
	ordered := make(map[int][]string)

	for col := 0; col < len(grid[0]); col++ {
		current := ""

		for row := range grid {
			if grid[row][col] != "" {
				current += grid[row][col]
			}
		}

		if len(current) == 0 {
			colIndex++
			continue
		}

		ordered[colIndex] = append(ordered[colIndex], current)
	}

	return ordered
}

func exectureOperations(ordered map[int][]string, operations []string) int {
	sum := 0

	for i, ops := range ordered {
		operation := operations[i]
		converted := utils.StringsToInts(ops)
		if operation == "+" {
			sum += utils.Sum(converted)
		} else {
			sum += utils.Product(converted)
		}
	}

	return sum
}

func Part2(input string) int {
	grid := utils.Parse(input, "")
	operations := extractOperations(grid[len(grid)-1])
	ordered := reorderNumbers(grid[:len(grid)-1])

	return exectureOperations(ordered, operations)
}
