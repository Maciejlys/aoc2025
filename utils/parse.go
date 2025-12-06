package utils

import (
	"strconv"
	"strings"
)

func LinesWithSeparator(fileContent string, separator string) []string {
	return strings.Split(strings.TrimRight(fileContent, "\n"), separator)
}

func Lines(fileContent string) []string {
	return strings.Split(strings.TrimRight(fileContent, "\n"), "\n")
}

func Parse(fileContent string, separator string) [][]string {
	lines := Lines(fileContent)
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

func StringsToInts(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, str := range strs {
		num := Atoi(str)
		ints = append(ints, num)
	}
	return ints
}
