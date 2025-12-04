package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolveDayTwo() {
	input := readInput()
	//result1 := calculateTotal(input)
	result2 := calculateTotal(input)

	fmt.Println("\n=== Day 2 ===")
	//fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}
func isInvalid(num int) bool {
	s := strconv.Itoa(num)
	length := len(s)
	if length%2 != 0 {
		return false
	}
	mid := length / 2
	return s[:mid] == s[mid:]
}

func isInvalid2(num int) bool {
	s := strconv.Itoa(num)
	length := len(s)
	if length <= 1 {
		return false
	}
	for l := 1; l <= length/2; l++ {
		if length%l == 0 {
			u := s[:l]
			isMatch := true
			for i := l; i < length; i += l {
				chunk := s[i : i+l]
				if chunk != u {
					isMatch = false
					break
				}
			}
			if isMatch {
				return true
			}
		}
	}
	return false
}

func calculateTotal(input string) int {
	total := 0
	ids := strings.SplitSeq(input, ",")

	for r := range ids {
		wholeId := strings.Split(strings.TrimSpace(r), "-")
		start, _ := strconv.Atoi(wholeId[0])
		end, _ := strconv.Atoi(wholeId[1])
		for num := start; num <= end; num++ {
			if isInvalid2(num) {
				total += num
			}
		}
	}
	return total
}

func readInput() string {
	content, err := os.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
