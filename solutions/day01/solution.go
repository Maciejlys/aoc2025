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
		if newPos <= 0 {
			d.MethodPassword += (utils.Abs(newPos) + d.Range) / d.Range

			if d.Position == 0 {
				d.MethodPassword--
			}
		}
		d.Position = (newPos + d.Range) % d.Range
	}

	utils.Log("Dial is rotated ", steps, " steps to the ", direction, " and now is at position ", d.Position, " Method password is ", d.MethodPassword)

	if d.Position == 0 {
		d.Password++
	}
}

func Part1(input string) int {
	dial := Dial{50, 100, 0, 0}
	for _, rotation := range utils.Lines(input) {
		var direction string
		var steps int
		fmt.Sscanf(rotation, "%1s%d", &direction, &steps)
		dial.Move(steps, direction)
	}
	return dial.Password
}

func Part2(input string, startPos int) int {
	dial := Dial{startPos, 100, 0, 0}

	utils.Log("Dial starts at ", dial.Position)

	for _, rotation := range utils.Lines(input) {
		var direction string
		var steps int
		fmt.Sscanf(rotation, "%1s%d", &direction, &steps)
		dial.Move(steps, direction)
	}

	utils.Log("Dial ends at ", dial.Position)

	return dial.MethodPassword
}
