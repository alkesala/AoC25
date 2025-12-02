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
	result1 := calculateTotal(input)
	fmt.Print(result1)
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

func calculateTotal(input string) int {
	total := 0
	ids := strings.SplitSeq(input, ",")

	for r := range ids {
		wholeId := strings.Split(strings.TrimSpace(r), "-")
		start, _ := strconv.Atoi(wholeId[0])
		end, _ := strconv.Atoi(wholeId[1])
		for num := start; num <= end; num++ {
			if isInvalid(num) {
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
