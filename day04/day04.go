package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

func part1(lines []string) {
	totalScore := 0

	for _, line := range lines {
		details := parseCardDetails(line)
		totalScore += details.score
	}

	fmt.Println(totalScore)
}

type CardDetails struct {
	have    []int
	winners map[int]bool
	score   int
	matches int
}

func parseCardDetails(line string) CardDetails {
	labelIndex := strings.Index(line, ":")
	split := strings.Split(line[labelIndex+1:], "|")
	haveStrings := strings.Split(strings.ReplaceAll(strings.TrimSpace(split[0]), "  ", " "), " ")
	winnerStrings := strings.Split(strings.ReplaceAll(strings.TrimSpace(split[1]), "  ", " "), " ")

	have := []int{}
	winners := map[int]bool{}
	matches := 0

	for _, str := range winnerStrings {
		parsed, _ := strconv.ParseInt(str, 10, 32)
		winners[int(parsed)] = true
	}

	for _, str := range haveStrings {
		parsed, _ := strconv.ParseInt(str, 10, 32)
		number := int(parsed)
		have = append(have, number)

		if winners[number] {
			matches++
		}
	}

	return CardDetails{
		have:    have,
		winners: winners,
		score:   getScoreForMatches(matches),
		matches: matches,
	}
}

func getScoreForMatches(matches int) int {
	if matches <= 1 {
		return matches
	}

	return util.IntPow(2, matches-1)
}

func part2(lines []string) {
	cardCounts := make([]int, len(lines))
	for i := range lines {
		cardCounts[i] = 1
	}

	for i, line := range lines {
		details := parseCardDetails(line)

		for j := 0; j < cardCounts[i]; j++ {
			for k := i + 1; k < i+1+details.matches && k < len(lines); k++ {
				cardCounts[k] = cardCounts[k] + 1
			}
		}
	}

	fmt.Println(util.SumSlice(cardCounts))
}

func Run() {
	lines := util.OpenAndRead("./day04/input.txt")

	part1(lines)
	part2(lines)
}
