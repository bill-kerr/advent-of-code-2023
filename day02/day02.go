package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

var cubeCounts map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(lines []string) {
	sum := 0

	for i, line := range lines {
		idx := strings.Index(line, ":")
		if isGamePossible(line[idx+2:]) {
			sum += i + 1
		}
	}

	fmt.Println(sum)
}

func isGamePossible(line string) bool {
	gameParts := strings.Split(line, "; ")
	for _, part := range gameParts {
		pulls := strings.Split(part, ", ")

		for _, pull := range pulls {
			pullDetails := strings.Split(pull, " ")
			num, _ := strconv.ParseInt(pullDetails[0], 10, 32)

			if int(num) > cubeCounts[pullDetails[1]] {
				return false
			}
		}
	}

	return true
}

func part2(lines []string) {
	sum := 0

	for _, line := range lines {
		idx := strings.Index(line, ":")
		sum += getPower(line[idx+2:])
	}

	fmt.Println(sum)
}

func getPower(line string) int {
	maxCounts := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	gameParts := strings.Split(line, "; ")
	for _, part := range gameParts {
		pulls := strings.Split(part, ", ")

		for _, pull := range pulls {
			pullDetails := strings.Split(pull, " ")
			num, _ := strconv.ParseInt(pullDetails[0], 10, 32)

			if int(num) > maxCounts[pullDetails[1]] {
				maxCounts[pullDetails[1]] = int(num)
			}
		}
	}

	return maxCounts["blue"] * maxCounts["red"] * maxCounts["green"]
}

func Run() {
	lines := util.OpenAndRead("./day02/input.txt")

	part1(lines)
	part2(lines)
}
