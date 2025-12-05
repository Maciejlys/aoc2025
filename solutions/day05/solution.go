package day05

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Maciejlys/aoc2025/utils"
)

type Range struct {
	min int
	max int
}

func (r Range) Contains(value int) bool {
	return value >= r.min && value <= r.max
}

func (r Range) Overlaps(other Range) bool {
	return r.min <= other.max && other.min <= r.max
}

func (r Range) Merge(other Range) Range {
	return Range{
		min: utils.Min(r.min, other.min),
		max: utils.Max(r.max, other.max),
	}
}

func mergeRanges(rs []Range) []Range {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].min < rs[j].min
	})

	merged := []Range{rs[0]}

	for _, curr := range rs[1:] {
		last := &merged[len(merged)-1]

		if last.Overlaps(curr) {
			*last = last.Merge(curr)
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}

type Ingredients struct {
	ranges []Range
	ids    []int
}

func extractIngredients(input string) Ingredients {
	ingredients := Ingredients{}
	parsed := strings.Split(strings.TrimSpace(input), "\n\n")
	ranges := []Range{}

	for _, r := range utils.Lines(parsed[0]) {
		var min int
		var max int
		fmt.Sscanf(r, "%d-%d", &min, &max)
		ranges = append(ranges, Range{min: min, max: max})
	}

	for _, id := range utils.Lines(parsed[1]) {
		var ticketID int
		fmt.Sscanf(id, "%d", &ticketID)
		ingredients.ids = append(ingredients.ids, ticketID)
	}

	ingredients.ranges = mergeRanges(ranges)

	return ingredients
}

func Part1(input string) int {
	ingredients := extractIngredients(input)
	sum := 0

	for _, id := range ingredients.ids {
		for _, r := range ingredients.ranges {
			if r.Contains(id) {
				sum++
				break
			}
		}
	}

	return sum
}

func Part2(input string) int {
	ingredients := extractIngredients(input)
	sum := 0

	for _, r := range ingredients.ranges {
		sum += r.max - r.min + 1
	}

	return sum
}
