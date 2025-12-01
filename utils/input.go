package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns all lines as a slice of strings
func ReadLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadFile reads the entire file content as a string
func ReadFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadInts reads a file and returns all lines as integers
func ReadInts(filepath string) ([]int, error) {
	lines, err := ReadLines(filepath)
	if err != nil {
		return nil, err
	}

	var nums []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

// SplitByEmptyLines splits input into groups separated by empty lines
func SplitByEmptyLines(lines []string) [][]string {
	var groups [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				groups = append(groups, current)
				current = []string{}
			}
		} else {
			current = append(current, line)
		}
	}

	if len(current) > 0 {
		groups = append(groups, current)
	}

	return groups
}

// ParseIntList parses a comma or space-separated list of integers
func ParseIntList(s string, separator string) ([]int, error) {
	parts := strings.Split(s, separator)
	var nums []int
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	
	return nums, nil
}
