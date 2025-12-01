package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadLines(t *testing.T) {
	// Create temp file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	
	content := "line1\nline2\nline3"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	lines, err := ReadLines(tmpFile)
	if err != nil {
		t.Fatalf("ReadLines() error = %v", err)
	}
	
	expected := []string{"line1", "line2", "line3"}
	if len(lines) != len(expected) {
		t.Errorf("ReadLines() returned %d lines, want %d", len(lines), len(expected))
	}
	
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Line %d = %q, want %q", i, line, expected[i])
		}
	}
}

func TestReadInts(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	
	content := "1\n2\n3\n"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	nums, err := ReadInts(tmpFile)
	if err != nil {
		t.Fatalf("ReadInts() error = %v", err)
	}
	
	expected := []int{1, 2, 3}
	if len(nums) != len(expected) {
		t.Errorf("ReadInts() returned %d numbers, want %d", len(nums), len(expected))
	}
	
	for i, num := range nums {
		if num != expected[i] {
			t.Errorf("Number %d = %d, want %d", i, num, expected[i])
		}
	}
}

func TestSplitByEmptyLines(t *testing.T) {
	lines := []string{
		"line1",
		"line2",
		"",
		"line3",
		"line4",
		"",
		"line5",
	}
	
	groups := SplitByEmptyLines(lines)
	
	if len(groups) != 3 {
		t.Errorf("SplitByEmptyLines() returned %d groups, want 3", len(groups))
	}
	
	if len(groups[0]) != 2 {
		t.Errorf("Group 0 has %d lines, want 2", len(groups[0]))
	}
	if len(groups[1]) != 2 {
		t.Errorf("Group 1 has %d lines, want 2", len(groups[1]))
	}
	if len(groups[2]) != 1 {
		t.Errorf("Group 2 has %d lines, want 1", len(groups[2]))
	}
}

func TestParseIntList(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		separator string
		expected  []int
	}{
		{"comma separated", "1,2,3", ",", []int{1, 2, 3}},
		{"space separated", "1 2 3", " ", []int{1, 2, 3}},
		{"with spaces", "1, 2, 3", ",", []int{1, 2, 3}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseIntList(tt.input, tt.separator)
			if err != nil {
				t.Fatalf("ParseIntList() error = %v", err)
			}
			
			if len(result) != len(tt.expected) {
				t.Errorf("ParseIntList() returned %d numbers, want %d", len(result), len(tt.expected))
			}
			
			for i, num := range result {
				if num != tt.expected[i] {
					t.Errorf("Number %d = %d, want %d", i, num, tt.expected[i])
				}
			}
		})
	}
}
