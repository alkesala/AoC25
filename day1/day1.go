package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolveDayOne() {
	input := readInput()
	result1 := dayOnePartOne(input)
	result2 := dayOnePartTwo(input)

	fmt.Println("\n=== Day 1 ===")
	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}
func dayOnePartOne(input string) int {

	var result int
	var start = 50
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			value, _ := strconv.Atoi(line[1:])
			value %= 100
			if line[0] == 'L' {
				value = -value
			}
			start += value
		}
		if start > 99 {
			start -= 100
		} else if start < 0 {
			start += 100
		}
		if start == 0 {
			result++
		}
	}
	return result
}

func dayOnePartTwo(input string) int {
	var result int
	var start = 50
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			value, _ := strconv.Atoi(line[1:])
			turns := value / 100
			result += turns
			value %= 100
			if line[0] == 'L' {
				value = -value
			}
			if start != 0 && value > 0 && start+value > 100 {
				result++
			}
			if start != 0 && value < 0 && start+value < 0 {
				result++
			}
			start = (start + value + 100) % 100
		}
		if start == 0 {
			result++
		}
	}
	return result
}

func readInput() string {
	content, err := os.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
