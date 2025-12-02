package day2

import (
	"testing"
)

func TestIsInvalid(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected bool
	}{
		{"single digit", 5, false},
		{"odd lenght", 123, false},
		{"odd lenght large", 12345, false},
		{"repeated single digit 22", 22, true},
		{"repeated single digit 66", 66, true},
		{"repeated two digit 1010", 1010, true},
		{"repeated two digit 2121", 2121, true},
		{"repeated two digit 9292", 9292, true},
		{"repeated four digit 14541454", 14541454, true},
		{"non-repeated four digit 1234", 1234, false},
		{"101 is valid", 101, false},
		{"1001 is valid", 1001, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isInvalid(tt.num)
			if result != tt.expected {
				t.Errorf("isInvalid(%d) = %v, want %v", tt.num, result, tt.expected)
			}
		})
	}

}

func TestCalculateTotal(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	expected := 1227775554

	result := calculateTotal(input)
	if result != expected {
		t.Errorf("calculateTotal() = %d, want %d", result, expected)
	}
}
