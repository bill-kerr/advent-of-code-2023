package day08

import (
	"fmt"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

type node struct {
	name  string
	left  string
	right string
}

func (n *node) isStart() bool {
	return rune(n.name[2]) == 'A'
}

func (n *node) isEnd() bool {
	return rune(n.name[2]) == 'Z'
}

func part1(lines []string) {
	instructions := lines[0]
	nodes := createNodes(lines[2:])

	currentNode := nodes["AAA"]
	currentInstruction := 0
	for currentNode.name != "ZZZ" {
		instruction := rune(instructions[currentInstruction%len(instructions)])

		if instruction == 'R' {
			currentNode = nodes[currentNode.right]
		} else {
			currentNode = nodes[currentNode.left]
		}
		currentInstruction++
	}

	fmt.Println(currentInstruction)
}

func part2(lines []string) {
	instructions := lines[0]
	nodes, positions := createNodesAndStartingPositions(lines[2:])
	endpoints := make([]int, len(positions))

	currentInstruction := 0
	for {
		instruction := rune(instructions[currentInstruction%len(instructions)])
		currentInstruction++

		for i := range positions {
			if instruction == 'R' {
				positions[i] = nodes[positions[i].right]
			} else {
				positions[i] = nodes[positions[i].left]
			}

			if positions[i].isEnd() && endpoints[i] == 0 {
				endpoints[i] = currentInstruction
			}
		}

		allEndpointsNonZero := true
		for _, endpoint := range endpoints {
			if endpoint == 0 {
				allEndpointsNonZero = false
				break
			}
		}

		if allEndpointsNonZero {
			break
		}
	}

	fmt.Println(util.LeastCommonMultiple(endpoints[0], endpoints[1], endpoints[2:]...))
}

func createNodes(lines []string) map[string]node {
	nodes := map[string]node{}

	for _, line := range lines {
		currentNodeName := line[:3]
		leftNodeName := line[7:10]
		rightNodeName := line[12:15]

		nodes[currentNodeName] = node{
			name:  currentNodeName,
			left:  leftNodeName,
			right: rightNodeName,
		}
	}

	return nodes
}

func createNodesAndStartingPositions(lines []string) (map[string]node, []node) {
	nodes := map[string]node{}
	startingPositions := []node{}

	for _, line := range lines {
		currentNodeName := line[:3]
		leftNodeName := line[7:10]
		rightNodeName := line[12:15]

		node := node{
			name:  currentNodeName,
			left:  leftNodeName,
			right: rightNodeName,
		}

		nodes[currentNodeName] = node

		if node.isStart() {
			startingPositions = append(startingPositions, node)
		}
	}

	return nodes, startingPositions
}

func Run() {
	lines := util.OpenAndRead("./day08/input.txt")
	part1(lines)
	part2(lines)
}
