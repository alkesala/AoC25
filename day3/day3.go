package day3

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	content, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
func SolveDayThree() {
	input := readInput()
	result1, result2 := findHighest(input)

	fmt.Println("\n=== Day 3 ===")
	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}

func findHighest(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	total := 0
	total12 := 0

	for _, line := range lines {
		bank := converter(line)
		maxJoltage := calcMaxJoltage(bank)
		total += maxJoltage
		maxJoltage12 := calcrMaxJoltage12(bank)
		total12 += maxJoltage12
	}
	return total, total12
}
func calcrMaxJoltage12(bank []int) int {
	keep := 12
	toBeRemoved := len(bank) - keep
	if toBeRemoved < 0 {
		keep = len(bank)
	}
	stack := []int{}
	for i := 0; i < len(bank); i++ {
		// pop the digits out of stack, feels dirty and unefficient but does the work
		for len(stack) > 0 && toBeRemoved > 0 && stack[len(stack)-1] < bank[i] {
			stack = stack[:len(stack)-1]
			toBeRemoved--
		}
		stack = append(stack, bank[i])
	}
	stack = stack[:keep]
	result := 0
	for _, digit := range stack {
		result = result*10 + digit
	}
	return result
}

func calcMaxJoltage(bank []int) int {
	first := 0
	for i := 0; i < len(bank)-1; i++ {
		if bank[i] > first {
			first = bank[i]
		}
	}
	// breaks for firstIndex
	firstIndex := -1
	for i := 0; i < len(bank)-1; i++ {
		if bank[i] == first {
			firstIndex = i
			break
		}
	}
	second := 0
	for i := firstIndex + 1; i < len(bank); i++ {
		if bank[i] > second {
			second = bank[i]
		}
	}
	// converts e.g., 1 and 9 to be 10 + 9 = 19
	return first*10 + second
}

func converter(line string) []int {
	digits := make([]int, 0, len(line))
	for _, char := range line {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal(err)
		}
		digits = append(digits, n)
	}
	return digits
}
