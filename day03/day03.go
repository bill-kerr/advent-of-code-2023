package day03

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

func part1(lines []string) {
	sum := 0

	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]

		j := 0
		for j < len(currentLine) {
			num, length, err := getNumAndLength(currentLine[j:])
			if err != nil {
				j++
				continue
			}

			var aboveLine *string
			var belowLine *string

			if i != 0 {
				aboveLine = &lines[i-1]
			}

			if i != len(lines)-1 {
				belowLine = &lines[i+1]
			}

			if isPartNumber(j, length, currentLine, aboveLine, belowLine) {
				sum += num
			}

			j += length
		}
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	sumOfRatios := 0

	for y, line := range lines {
		for x, char := range line {
			if char == '*' {
				numbers := findAdjacentNumbers(lines, x, y)
				if len(numbers) == 2 {
					sumOfRatios += numbers[0] * numbers[1]
				}
			}
		}
	}

	fmt.Println(sumOfRatios)
}

func findAdjacentNumbers(lines []string, x, y int) []int {
	numbers := []int{}
	currentLine := &lines[y]

	right, err := getNumberAt(currentLine, x+1)
	if err == nil {
		numbers = append(numbers, right)
	}

	left, err := getNumberAt(currentLine, x-1)
	if err == nil {
		numbers = append(numbers, left)
	}

	aboveLine := getLineAt(lines, y-1)
	belowLine := getLineAt(lines, y+1)

	if aboveLine != nil {
		numbers = append(numbers, findAdjacentNumbersOnLine(*aboveLine, x)...)
	}

	if belowLine != nil {
		numbers = append(numbers, findAdjacentNumbersOnLine(*belowLine, x)...)
	}

	return numbers
}

func findAdjacentNumbersOnLine(line string, x int) []int {
	numbers := []int{}

	foundNumber := false
	position := x - 1
	for position >= 0 && position <= x+1 && position <= len(line)-1 {
		num, err := getNumberAt(&line, position)
		if err != nil {
			foundNumber = false
			position++
			continue
		}

		if !foundNumber {
			numbers = append(numbers, num)
			foundNumber = true
			position++
			continue
		}

		position++
	}

	return numbers
}

func getLineAt(lines []string, y int) *string {
	if y < 0 || y > len(lines)-1 {
		return nil
	}
	return &(lines[y])
}

func getNumberAt(line *string, index int) (int, error) {
	numAsString := ""
	startDigit, err := util.RtoDigit(rune((*line)[index]))
	if err != nil {
		return 0, errors.New("not a number")
	}

	numAsString += fmt.Sprint(startDigit)

	i := index + 1
	j := index - 1

	// Search right
	for i <= len(*line)-1 {
		rightDigit, err := util.RtoDigit(rune((*line)[i]))
		if err != nil {
			break
		}
		numAsString += fmt.Sprint(rightDigit)
		i++
	}

	// Search left
	for j >= 0 {
		leftDigit, err := util.RtoDigit(rune((*line)[j]))
		if err != nil {
			break
		}
		numAsString = fmt.Sprintf("%v%v", leftDigit, numAsString)
		j--
	}

	parsed, err := strconv.ParseInt(numAsString, 10, 32)
	return int(parsed), err
}

func isDigit(char rune) bool {
	_, err := util.RtoDigit(char)
	return err == nil
}

func isSymbol(char rune) bool {
	return !isDigit(char) && char != '.'
}

func getNumAndLength(line string) (int, int, error) {
	numAsSlice := []rune{}

	for _, char := range line {
		if isDigit(char) {
			numAsSlice = append(numAsSlice, char)
		} else if len(numAsSlice) > 0 {
			parsed, err := strconv.ParseInt(string(numAsSlice), 10, 32)
			if err == nil {
				return int(parsed), len(numAsSlice), nil
			}

			return 0, 0, errors.New("failed to parse integer")
		} else {
			return 0, 0, errors.New("no number")
		}
	}

	parsed, err := strconv.ParseInt(string(numAsSlice), 10, 32)
	if err == nil {
		return int(parsed), len(numAsSlice), nil
	}

	return 0, 0, errors.New("no number")
}

func isPartNumber(index int, length int, currentLine string, aboveLine *string, belowLine *string) bool {
	if index != 0 && rune(currentLine[index-1]) != '.' {
		return true
	}

	if index+length <= len(currentLine)-1 && rune(currentLine[index+length]) != '.' {
		return true
	}

	for _, line := range []*string{aboveLine, belowLine} {
		if line == nil {
			continue
		}

		start := math.Max(0, float64(index-1))
		end := math.Min(float64(len(currentLine)), float64(index+length+1))
		sliced := (*line)[int(start):int(end)]

		for _, char := range sliced {
			if isSymbol(char) {
				return true
			}
		}
	}

	return false
}

func Run() {
	lines := util.OpenAndRead("./day03/input.txt")

	part1(lines)
	part2(lines)
}
