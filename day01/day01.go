package day01

import (
	"errors"
	"fmt"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

func part1(lines []string) {
	sum := 0

	for _, line := range lines {
		var first *int
		var last *int

		for _, char := range line {
			digit, err := util.RtoDigit(char)
			if err != nil {
				continue
			}

			if first == nil {
				first = &digit
			}

			last = &digit
		}

		sum += *first*10 + *last
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	sum := 0

	for _, line := range lines {
		var first *int
		var last *int

		for i := range line {
			digit, err := getDigit([]rune(line), i)
			if err != nil {
				continue
			}

			if first == nil {
				first = &digit
			}

			last = &digit
		}

		sum += *first*10 + *last
	}

	fmt.Println(sum)
}

func getDigit(runes []rune, currentIndex int) (int, error) {
	digit, err := util.RtoDigit(runes[currentIndex])
	if err == nil {
		return digit, nil
	}

	word := ""
	for i := currentIndex; i < len(runes); i++ {
		word += string(runes[i])
		digitFromWord, err := getDigitFromWord(word)
		if err == nil {
			return digitFromWord, nil
		}
	}

	return 0, errors.New("no digit")
}

var digitValuesMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getDigitFromWord(str string) (int, error) {
	digit, ok := digitValuesMap[str]
	if ok {
		return digit, nil
	}

	return 0, errors.New("not a digit")
}

func Run() {
	lines := util.OpenAndRead("./day01/input.txt")

	part1(lines)
	part2(lines)
}
