package day04

import (
	"os"
	"strings"
)

type cell struct {
	number int
	draw   bool
}

func read_input() ([]int, [][][]cell) {
	input, _ := os.ReadFile("../inputs/day04.txt")
	lines := strings.Split(string(input), "\n")
	nums := lines[0]

	tables := strings.Split(string(input), "\n\n")
	tables = tables[1:]
	for table := 0; table < len(tables); table++ {
		tables[table] := strings.Split(string(tables[table]), "\n")

	}
	return nums, tables
}
