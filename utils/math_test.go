package utils

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
	}
	
	for _, tt := range tests {
		result := Abs(tt.input)
		if result != tt.expected {
			t.Errorf("Abs(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 1},
		{5, 3, 3},
		{-1, -5, -5},
	}
	
	for _, tt := range tests {
		result := Min(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Min(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{-1, -5, -1},
	}
	
	for _, tt := range tests {
		result := Max(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{-1, 1}, 0},
		{[]int{}, 0},
		{[]int{10}, 10},
	}
	
	for _, tt := range tests {
		result := Sum(tt.input)
		if result != tt.expected {
			t.Errorf("Sum(%v) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{2, 5}, 10},
		{[]int{}, 0},
		{[]int{10}, 10},
	}
	
	for _, tt := range tests {
		result := Product(tt.input)
		if result != tt.expected {
			t.Errorf("Product(%v) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{12, 8, 4},
		{15, 5, 5},
		{7, 13, 1},
	}
	
	for _, tt := range tests {
		result := GCD(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{12, 8, 24},
		{15, 5, 15},
		{7, 13, 91},
	}
	
	for _, tt := range tests {
		result := LCM(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("LCM(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
