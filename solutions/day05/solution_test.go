package day05

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
			want:  862,
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
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  14,
		},
		{
			name:  "input",
			input: input,
			want:  357907198933892,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func BenchmarkPart1(b *testing.B) {
// 	for b.Loop() {
// 		Part1(input)
// 	}
// }
//
// func BenchmarkPart2(b *testing.B) {
// 	for b.Loop() {
// 		Part2(input)
// 	}
// }
