package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Maciejlys/aoc2025/day01"
)

func main() {
	dayFlag := flag.Int("day", 1, "Which day to run (1-25)")
	inputFlag := flag.String("input", "", "Path to input file (optional, defaults to dayXX/input.txt)")
	flag.Parse()

	day := *dayFlag
	inputPath := *inputFlag

	// If no input path specified, use default
	if inputPath == "" {
		inputPath = fmt.Sprintf("day%02d/input.txt", day)
	}

	// Check if input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Printf("Warning: Input file %s not found. You may need to download it from https://adventofcode.com/2025/day/%d/input\n\n", inputPath, day)
	}

	// Run the specified day
	switch day {
	case 1:
		day01.Run(inputPath)
	// Add more days here as you implement them
	// case 2:
	//     day02.Run(inputPath)
	default:
		fmt.Printf("Day %d is not implemented yet\n", day)
		os.Exit(1)
	}
}
