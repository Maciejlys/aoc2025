package day1

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  3,
		},
		{
			name:  "input",
			input: input,
			want:  1139,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		startPos int
		want     int
	}{
		{
			name:     "example",
			input:    example,
			startPos: 50,
			want:     6,
		},
		{
			name:     "example2",
			startPos: 5,
			input:    "L10",
			want:     1,
		},
		{
			name:     "example3",
			startPos: 95,
			input:    "R10",
			want:     1,
		},
		{
			name:     "example4",
			startPos: 0,
			input:    "R10",
			want:     0,
		},
		{
			name:     "example5",
			startPos: 0,
			input:    "L10",
			want:     0,
		},
		{
			name:     "example5",
			startPos: 10,
			input:    "L10",
			want:     1,
		},
		{
			name:     "example6",
			startPos: 90,
			input:    "R10",
			want:     1,
		},
		{
			name:     "example7",
			startPos: 10,
			input:    "L110",
			want:     2,
		},
		{
			name:     "example8",
			startPos: 10,
			input:    "R190",
			want:     2,
		},
		{
			name:     "example9",
			startPos: 5,
			input:    "L6",
			want:     1,
		},
		{
			name:     "example10",
			startPos: 0,
			input:    "L100",
			want:     1,
		},
		{
			name:     "example10",
			startPos: 0,
			input:    "R100",
			want:     1,
		},
		{
			name:     "example11",
			startPos: 50,
			input:    "R1000",
			want:     10,
		},
		{
			name:     "input",
			input:    input,
			startPos: 50,
			want:     6684,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.input, tt.startPos); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
