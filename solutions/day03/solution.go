package day03

import (
	"strconv"
	"strings"

	"github.com/Maciejlys/aoc2025/utils"
)

func FindJoinedBattery(numbers []int, k int) int {
	remove := len(numbers) - k
	stack := utils.NewStack[int]()

	for _, num := range numbers {
		for remove > 0 && !stack.IsEmpty() && (stack.Top() < num) {
			stack.Pop()
			remove--
		}
		stack.Push(num)
	}

	highest := stack.ToSlice()[:k]

	var sb strings.Builder
	for _, d := range highest {
		sb.WriteByte(byte(d) + '0')
	}

	joinedStr := sb.String()

	joinedInt, _ := strconv.Atoi(joinedStr)

	return joinedInt
}

func Part1(input string) int {
	parsed := utils.ParseWithInts(input, "")
	sum := 0

	for _, line := range parsed {
		sum += FindJoinedBattery(line, 2)
	}

	return sum
}

func Part2(input string) int {
	parsed := utils.ParseWithInts(input, "")
	sum := 0

	for _, line := range parsed {
		sum += FindJoinedBattery(line, 12)
	}

	return sum
}
