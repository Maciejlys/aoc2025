package day1

import (
	"fmt"

	"github.com/Maciejlys/aoc2025/utils"
)

type Dial struct {
	Position       int
	Range          int
	Password       int
	MethodPassword int
}

func (d *Dial) Move(steps int, direction string) {
	switch direction {
	case "R":
		newPos := d.Position + steps
		if newPos >= d.Range {
			d.MethodPassword += newPos / d.Range
		}
		d.Position = newPos % d.Range
	case "L":
		newPos := d.Position - steps

		if newPos < 0 {
			if d.Position == 0 {
				d.MethodPassword += utils.Abs(newPos) / d.Range
			} else {
				remainingSteps := utils.Abs(newPos)
				d.MethodPassword += (remainingSteps / d.Range) + 1
			}
		} else if newPos == 0 && d.Position > 0 {
			d.MethodPassword += 1
		}

		d.Position = ((newPos % d.Range) + d.Range) % d.Range
	}

	if d.Position == 0 {
		d.Password++
	}
}

func moveDial(dial *Dial, input string) {
	for _, rotation := range utils.Lines(input) {
		var direction string
		var steps int
		fmt.Sscanf(rotation, "%1s%d", &direction, &steps)
		dial.Move(steps, direction)
	}
}

func Part1(input string) int {
	dial := Dial{50, 100, 0, 0}
	moveDial(&dial, input)
	return dial.Password
}

func Part2(input string, startPos int) int {
	dial := Dial{startPos, 100, 0, 0}
	moveDial(&dial, input)
	return dial.MethodPassword
}
