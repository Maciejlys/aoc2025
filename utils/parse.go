package utils

import (
	"strconv"
	"strings"
)

func Parse(fileContent string, separator string) [][]string {
	lines := strings.Split(strings.TrimRight(fileContent, "\n"), "\n")
	result := make([][]string, 0)

	for _, line := range lines {
		parts := strings.Split(line, separator)
		for i, part := range parts {
			parts[i] = strings.TrimSpace(part)
		}

		result = append(result, parts)
	}

	return result
}

func ParseWithInts(fileContent string, separator string) [][]int {
	parsed := Parse(fileContent, separator)
	result := make([][]int, 0)
	for _, line := range parsed {
		intLine := make([]int, 0)
		for _, strNum := range line {
			intNum, _ := strconv.Atoi(strNum)
			intLine = append(intLine, intNum)
		}
		result = append(result, intLine)
	}

	return result
}
