package day02

import (
	"fmt"
	"strings"

	"github.com/Maciejlys/aoc2025/utils"
)

func IsInvalidId(id int) bool {
	strId := fmt.Sprintf("%d", id)

	if len(strId)%2 != 0 {
		return false
	}

	mid := len(strId) / 2

	return strId[:mid] == strId[mid:]
}

func IsInvalidIdP2(id int) bool {
	strId := fmt.Sprintf("%d", id)
	mid := len(strId) / 2

	for i := range mid {
		pattern := strId[:i+1]
		repeat := strings.Repeat(pattern, len(strId)/len(pattern))

		if repeat == strId {
			return true
		}
	}

	return false
}

func extractInvalidIds(start int, end int) []int {
	invalidIds := make([]int, 0)

	for i := start; i <= end; i++ {
		if IsInvalidId(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func Part1(input string) int {
	parsed := utils.LinesWithSeparator(input, ",")
	sum := 0

	for _, id := range parsed {
		var start int
		var end int
		fmt.Sscanf(id, "%d-%d", &start, &end)
		invalidIds := extractInvalidIds(start, end)
		sum += utils.Sum(invalidIds)
	}

	return sum
}

func extractInvalidIdsP2(start int, end int) []int {
	invalidIds := make([]int, 0)

	for i := start; i <= end; i++ {
		if IsInvalidIdP2(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func Part2(input string) int {
	parsed := utils.LinesWithSeparator(input, ",")
	sum := 0

	for _, id := range parsed {
		var start int
		var end int
		fmt.Sscanf(id, "%d-%d", &start, &end)
		invalidIds := extractInvalidIdsP2(start, end)
		sum += utils.Sum(invalidIds)
	}

	return sum
}
