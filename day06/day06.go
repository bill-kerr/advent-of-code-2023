package day06

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

func part1(lines []string) {
	times := getNumbers(lines[0])
	distances := getNumbers(lines[1])
	product := 1

	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]

		// x = y*(z-y)    Where x is distance, y is time spent holding button, and z is total time of race
		// Therefore, there are two points we are interested in: the two points where x = distance
		// For instance, 9 = y*(7-y) ==> y^2 - 7y - 9 = 0
		// Generalized, it's (y ± sqrt(y^2-4x))/2
		pointA, pointB := solveQuadratic(time, distance)

		a := int(math.Ceil(pointA - 1))
		b := int(math.Floor(pointB + 1))

		product *= a - b + 1
	}

	fmt.Println(product)
}

func part2(lines []string) {
	time := getNumber(lines[0])
	distance := getNumber(lines[1])

	pointA, pointB := solveQuadratic(time, distance)

	a := int(math.Ceil(pointA - 1))
	b := int(math.Floor(pointB + 1))

	fmt.Println(a - b + 1)
}

func getNumbers(line string) []int {
	return util.ParseInts(strings.Fields(line)[1:])
}

func getNumber(line string) int {
	parsed, _ := strconv.ParseInt(strings.Join(strings.Fields(line)[1:], ""), 10, 64)
	return int(parsed)
}

func solveQuadratic(b int, c int) (float64, float64) {
	solutionA := (float64(b) + math.Sqrt(float64(math.Pow(float64(b), 2))-float64(4*c))) / 2
	solutionB := (float64(b) - math.Sqrt(float64(math.Pow(float64(b), 2))-float64(4*c))) / 2
	return solutionA, solutionB
}

func Run() {
	lines := util.OpenAndRead("./day06/input.txt")

	part1(lines)
	part2(lines)
}
