package day09

import (
	"fmt"
	"strings"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

func part1(lines []string) {
	sum := 0

	for _, line := range lines {
		numbers := util.ParseInts(strings.Fields(line))
		lastNumbers := []int{}

		for {
			lastNumber := numbers[len(numbers)-1]
			lastNumbers = append(lastNumbers, lastNumber)
			if util.Every(numbers, func(number int, _ int) bool {
				return number == lastNumber
			}) {
				break
			}

			numbers = takeDifferences(numbers)
		}

		sum += util.SumSlice(lastNumbers)
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	sum := 0

	for _, line := range lines {
		numbers := util.ParseInts(strings.Fields(line))
		firstNumbers := []int{}

		for {
			firstNumber := numbers[0]
			firstNumbers = append(firstNumbers, firstNumber)
			if util.Every(numbers, func(number int, _ int) bool {
				return number == firstNumber
			}) {
				break
			}

			numbers = takeDifferences(numbers)
		}

		prediction := 0
		for i := len(firstNumbers) - 1; i >= 0; i-- {
			prediction = firstNumbers[i] - prediction
		}
		sum += prediction
	}

	fmt.Println(sum)
}

func takeDifferences(numbers []int) []int {
	differences := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		differences[i] = numbers[i+1] - numbers[i]
	}
	return differences
}

func Run() {
	lines := util.OpenAndRead("./day09/input.txt")

	part1(lines)
	part2(lines)
}
