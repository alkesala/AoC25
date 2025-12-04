package day3

import (
	"testing"
)

func TestFindHighest(t *testing.T) {
	tests := []struct {
		name     string
		num      string
		expected int
	}{
		{"digits 987654321111111", "987654321111111", 98},
		{"digits 811111111111119", "811111111111119", 89},
		{"digits 234234234234278", "234234234234278", 78},
		{"digits 818181911112111", "18181911112111", 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := findHighest(tt.num)
			if result != tt.expected {
				t.Errorf("isInvalid(%s) = %v, want %v", tt.num, result, tt.expected)
			}
		})
	}
}

func TestFindHighestMaxJoltage(t *testing.T) {
	tests := []struct {
		name     string
		num      string
		expected int
	}{
		{"digits 987654321111111", "987654321111111", 987654321111},
		{"digits 811111111111119", "811111111111119", 811111111119},
		{"digits 234234234234278", "234234234234278", 434234234278},
		{"digits 818181911112111", "818181911112111", 888911112111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result := findHighest(tt.num)
			if result != tt.expected {
				t.Errorf("isInvalid(%s) = %v, want %v", tt.num, result, tt.expected)
			}
		})
	}
}
