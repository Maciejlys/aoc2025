package day01

import (
	"fmt"

	"github.com/Maciejlys/aoc2025/utils"
)

// Part1 solves part 1 of the puzzle
func Part1(inputPath string) (int, error) {
	lines, err := utils.ReadLines(inputPath)
	if err != nil {
		return 0, err
	}

	// TODO: Implement solution
	_ = lines
	
	return 0, nil
}

// Part2 solves part 2 of the puzzle
func Part2(inputPath string) (int, error) {
	lines, err := utils.ReadLines(inputPath)
	if err != nil {
		return 0, err
	}

	// TODO: Implement solution
	_ = lines
	
	return 0, nil
}

// Run executes both parts and prints results
func Run(inputPath string) {
	fmt.Println("Day 01:")
	
	result1, err := Part1(inputPath)
	if err != nil {
		fmt.Printf("  Part 1 Error: %v\n", err)
	} else {
		fmt.Printf("  Part 1: %d\n", result1)
	}
	
	result2, err := Part2(inputPath)
	if err != nil {
		fmt.Printf("  Part 2 Error: %v\n", err)
	} else {
		fmt.Printf("  Part 2: %d\n", result2)
	}
}
