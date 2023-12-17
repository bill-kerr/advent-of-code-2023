package main

import (
	"os"

	"github.com/bill-kerr/advent-of-code-2023/day10"
	"github.com/bill-kerr/advent-of-code-2023/util"
)

func main() {
	if len(os.Args) >= 3 {
		t := os.Args[1]
		name := os.Args[2]

		if t == "-t" && name != "" {
			util.CreateTemplate(name)
		}
	} else {
		day10.Run()
	}
}
