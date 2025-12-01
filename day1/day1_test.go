package day1

import "testing"

func TestDayOnePartOne(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	result := dayOnePartOne(input)

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	result := dayOnePartTwo(input)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
