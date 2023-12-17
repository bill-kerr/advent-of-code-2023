package day10

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

type pipe rune

type position struct {
	x, y int
}

func (p *position) equals(other position) bool {
	return p.x == other.x && p.y == other.y
}

func (p *position) move(direction direction) {
	switch direction {
	case left:
		p.x -= 1
	case right:
		p.x += 1
	case up:
		p.y -= 1
	case down:
		p.y += 1
	}
}

type direction int

func (d direction) opposite() direction {
	switch d {
	case up:
		return down
	case down:
		return up
	case right:
		return left
	case left:
		return right
	default:
		return undefined
	}
}

type connection struct {
	direction direction
	pipe      pipe
}

const (
	up direction = iota
	right
	down
	left
	undefined
)

const (
	vertical      pipe = '|' // up, down
	horizontal    pipe = '-' // right, left
	bendNorthEast pipe = 'L' // right, up
	bendNorthWest pipe = 'J' // left, up
	bendSouthWest pipe = '7' // left, down
	bendSouthEast pipe = 'F' // right, down
	start         pipe = 'S'
	none          pipe = '.'
)

func (p pipe) connects(direction direction) bool {
	switch direction {
	case up:
		return p == bendNorthWest || p == bendNorthEast || p == vertical
	case down:
		return p == bendSouthWest || p == bendSouthEast || p == vertical
	case left:
		return p == bendNorthWest || p == bendSouthWest || p == horizontal
	case right:
		return p == bendNorthEast || p == bendSouthEast || p == horizontal
	default:
		return false
	}
}

func part1(lines []string) {
	pipeMap, startPosition := parseMap(lines)

	initialConnections := getConnections(startPosition, pipeMap)
	directionA, directionB := selectInitialDirections(initialConnections)

	positionA := startPosition
	positionA.move(directionA)

	positionB := startPosition
	positionB.move(directionB)

	steps := 0
	for !positionA.equals(positionB) {
		steps++

		directionA = selectNextDirection(directionA.opposite(), getConnections(positionA, pipeMap))
		positionA.move(directionA)

		directionB = selectNextDirection(directionB.opposite(), getConnections(positionB, pipeMap))
		positionB.move(directionB)
	}

	fmt.Println(steps + 1)
}

func part2(lines []string) {
	pipeMap, startPosition := parseMap(lines)

	initialConnections := getConnections(startPosition, pipeMap)
	directionA, directionB := selectInitialDirections(initialConnections)
	pipeMap[startPosition.y][startPosition.x] = findStartPipeType(initialConnections)

	positionA := startPosition
	positionA.move(directionA)

	positionB := startPosition
	positionB.move(directionB)

	pipeCoords := make([][]bool, len(lines))
	for i := 0; i < len(lines); i++ {
		pipeCoords[i] = make([]bool, len(lines[i]))
	}

	pipeCoords[positionA.y][positionA.x] = true
	pipeCoords[positionB.y][positionB.x] = true
	pipeCoords[startPosition.y][startPosition.x] = true

	steps := 0
	for !positionA.equals(positionB) {
		steps++

		directionA = selectNextDirection(directionA.opposite(), getConnections(positionA, pipeMap))
		positionA.move(directionA)

		directionB = selectNextDirection(directionB.opposite(), getConnections(positionB, pipeMap))
		positionB.move(directionB)

		pipeCoords[positionA.y][positionA.x] = true
		pipeCoords[positionB.y][positionB.x] = true
	}

	// Now we use ray casting to determine points within the enclosed pipe network
	isWithin := false
	emptyCount := 0
	for y := 0; y < len(lines); y++ {
		isWithin = false
		previousPipe := none

		for x := 0; x < len(lines[y]); x++ {
			if pipeCoords[y][x] && pipeMap[y][x] == vertical {
				isWithin = !isWithin
			}

			if pipeCoords[y][x] && pipeMap[y][x] == bendNorthWest && previousPipe == bendSouthEast {
				isWithin = !isWithin
				previousPipe = none
			} else if pipeCoords[y][x] && pipeMap[y][x] == bendSouthWest && previousPipe == bendNorthEast {
				isWithin = !isWithin
				previousPipe = none
			}

			if pipeCoords[y][x] && pipeMap[y][x] == bendSouthEast {
				previousPipe = bendSouthEast
			} else if pipeCoords[y][x] && pipeMap[y][x] == bendNorthEast {
				previousPipe = bendNorthEast
			}

			if isWithin && !pipeCoords[y][x] {
				emptyCount++
			}
		}
	}

	fmt.Println(emptyCount)
}

func parseMap(lines []string) ([][]pipe, position) {
	pipes := make([][]pipe, len(lines))
	startPosition := position{}

	for i, line := range lines {
		pipeLine := make([]pipe, len(line))
		for j, char := range line {
			pipeLine[j] = pipe(char)

			if pipeLine[j] == start {
				startPosition.x = j
				startPosition.y = i
			}
		}

		pipes[i] = pipeLine
	}

	return pipes, startPosition
}

func findStartPipeType(connections map[direction]connection) pipe {
	_, connectsUp := connections[up]
	_, connectsDown := connections[down]
	_, connectsLeft := connections[left]
	_, connectsRight := connections[right]

	if connectsUp && connectsDown {
		return vertical
	}

	if connectsLeft && connectsRight {
		return horizontal
	}

	if connectsLeft && connectsDown {
		return bendSouthWest
	}

	if connectsLeft && connectsUp {
		return bendNorthWest
	}

	if connectsRight && connectsDown {
		return bendSouthEast
	}

	if connectsRight && connectsUp {
		return bendNorthEast
	}

	return none
}

func getConnections(position position, pipes [][]pipe) map[direction]connection {
	if pipes[position.y][position.x] == start {
		return getStartingConnections(position, pipes)
	}

	connections := map[direction]connection{}

	if position.x > 0 && pipes[position.y][position.x].connects(left) {
		connections[left] = connection{direction: left, pipe: pipes[position.y][position.x-1]}
	}

	if position.x < len(pipes[position.y])-1 && pipes[position.y][position.x].connects(right) {
		connections[right] = connection{direction: right, pipe: pipes[position.y][position.x+1]}
	}

	if position.y > 0 && pipes[position.y][position.x].connects(up) {
		connections[up] = connection{direction: up, pipe: pipes[position.y-1][position.x]}
	}

	if position.y < len(pipes)-1 && pipes[position.y][position.x].connects(down) {
		connections[down] = connection{direction: down, pipe: pipes[position.y+1][position.x]}
	}

	return connections
}

func getStartingConnections(position position, pipes [][]pipe) map[direction]connection {
	connections := map[direction]connection{}

	if position.x > 0 && pipes[position.y][position.x-1].connects(right) {
		connections[left] = connection{direction: left, pipe: pipes[position.y][position.x-1]}
	}

	if position.x < len(pipes[position.y])-1 && pipes[position.y][position.x+1].connects(left) {
		connections[right] = connection{direction: right, pipe: pipes[position.y][position.x+1]}
	}

	if position.y > 0 && pipes[position.y-1][position.x].connects(down) {
		connections[up] = connection{direction: up, pipe: pipes[position.y-1][position.x]}
	}

	if position.y < len(pipes)-1 && pipes[position.y+1][position.x].connects(up) {
		connections[down] = connection{direction: down, pipe: pipes[position.y+1][position.x]}
	}

	return connections
}

func selectInitialDirections(connections map[direction]connection) (direction, direction) {
	aDirection := undefined
	bDirection := undefined

	for i := 0; i < 4; i++ {
		_, aConnects := connections[direction(i)]
		_, bConnects := connections[direction(3-i)]

		if aDirection == undefined && aConnects {
			aDirection = direction(i)
		}

		if bDirection == undefined && bConnects {
			bDirection = direction(3 - i)
		}
	}

	return aDirection, bDirection
}

func selectNextDirection(previousDirection direction, connections map[direction]connection) direction {
	for i := 0; i < 4; i++ {
		_, connects := connections[direction(i)]
		if connects && direction(i) != previousDirection {
			return direction(i)
		}
	}

	fmt.Println(previousDirection, connections)
	panic("no possible next direction!")
}

func Run() {
	lines := util.OpenAndRead("./day10/input.txt")

	part1(lines)
	part2(lines)
}
