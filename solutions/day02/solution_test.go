package day02

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func Test_IsInalidId(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "11",
			input: 11,
			want:  true,
		},
		{
			name:  "22",
			input: 22,
			want:  true,
		},
		{
			name:  "21",
			input: 21,
			want:  false,
		},
		{
			name:  "1010",
			input: 1010,
			want:  true,
		},
		{
			name:  "38593859",
			input: 38593859,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInvalidId(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  1227775554,
		},
		{
			name:  "input",
			input: input,
			want:  28846518423,
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

func Test_IsInalidIdP2(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "21",
			input: 21,
			want:  false,
		},
		{
			name:  "2115",
			input: 2115,
			want:  false,
		},
		{
			name:  "11",
			input: 11,
			want:  true,
		},
		{
			name:  "22",
			input: 22,
			want:  true,
		},
		{
			name:  "999",
			input: 999,
			want:  true,
		},
		{
			name:  "1010",
			input: 1010,
			want:  true,
		},
		{
			name:  "222222",
			input: 222222,
			want:  true,
		},
		{
			name:  "565656",
			input: 565656,
			want:  true,
		},
		{
			name:  "824824824",
			input: 824824824,
			want:  true,
		},
		{
			name:  "2121212121",
			input: 2121212121,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInvalidIdP2(tt.input); got != tt.want {
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
			want:  4174379265,
		},
		{
			name:  "input",
			input: input,
			want:  31578210022,
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

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		Part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		Part2(input)
	}
}
