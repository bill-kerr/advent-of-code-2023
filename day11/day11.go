package day11

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

type position struct {
	x int
	y int
}

func part1(lines []string) {
	galaxies := findExpandedGalaxyPositions(lines, 2)
	distanceSum := sumGalaxyDistances(galaxies)
	fmt.Println(distanceSum)
}

func part2(lines []string) {
	galaxies := findExpandedGalaxyPositions(lines, 1000000)
	distanceSum := sumGalaxyDistances(galaxies)
	fmt.Println(distanceSum)
}

func getEmptyRowsAndColumns(lines []string) (map[int]bool, map[int]bool) {
	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}
	nonEmptyCols := map[int]bool{}

	for i, line := range lines {
		isRowEmpty := true

		for j, char := range line {
			if char == '#' {
				isRowEmpty = false
				nonEmptyCols[j] = true
			}

			if i == len(lines)-1 && !nonEmptyCols[j] {
				emptyCols[j] = true
			}
		}

		if isRowEmpty {
			emptyRows[i] = true
		}
	}

	return emptyRows, emptyCols
}

func findExpandedGalaxyPositions(lines []string, expansionFactor int) []position {
	emptyRows, emptyCols := getEmptyRowsAndColumns(lines)
	positions := []position{}

	yOffset := 0

	for y, line := range lines {
		xOffset := 0

		for x, char := range line {
			if char == '#' {
				positions = append(positions, position{x: xOffset, y: yOffset})
			}

			if emptyCols[x] {
				xOffset += expansionFactor
			} else {
				xOffset++
			}
		}

		if emptyRows[y] {
			yOffset += expansionFactor
		} else {
			yOffset++
		}
	}

	return positions
}

func sumGalaxyDistances(galaxies []position) int {
	sum := 0

	for i, origin := range galaxies {
		for _, destination := range galaxies[i+1:] {
			xDistance := util.AbsInt(origin.x - destination.x)
			yDistance := util.AbsInt(origin.y - destination.y)
			sum += xDistance + yDistance
		}
	}

	return sum
}

func Run() {
	lines := util.OpenAndRead("./day11/input.txt")

	part1(lines)
	part2(lines)
}
