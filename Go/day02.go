package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type course struct {
	way  string
	dist int
}

func read_input() (input []course) {
	b, _ := os.ReadFile("../inputs/day02.txt")

	lines := strings.Split(string(b), "\n")

	//	input := make([]course, 0, len(lines))

	for _, l := range lines {

		if len(l) == 0 {
			continue
		}

		e := strings.Split(l, " ")
		n, _ := strconv.Atoi(e[1])

		input = append(input, course{way: e[0], dist: n})
	}

	return input
}

func day02_I() {
	//test := []course{{way: "forward", dist: 5}, {"down", 5}, {"forward", 8}, {"up", 3}, {"down", 8}, {"forward", 2}}
	input := read_input()
	x := 0
	y := 0
	for i := 0; i < len(input); i++ {
		if input[i].way == "forward" {
			x = x + input[i].dist
		}
		if input[i].way == "down" {
			y += input[i].dist
		}
		if input[i].way == "up" {
			y -= input[i].dist
		}
	}
	println("X : ", x, "\nY : ", y)
	fmt.Println("Ans : ", x*y)
}

func day02_II() {
	input := read_input()
	x := 0
	y := 0
	aim := 0
	for i := 0; i < len(input); i++ {
		if input[i].way == "forward" {
			x += input[i].dist
			y += aim * input[i].dist
		}
		if input[i].way == "down" {
			aim += input[i].dist
		}
		if input[i].way == "up" {
			aim -= input[i].dist
		}
	}
	println("X : ", x, "\nY : ", y)
	fmt.Println("Ans : ", x*y)
}

func main() {
	fmt.Println("*** ADVENT OF CODE 2021 ** DAY02 ***")
	fmt.Println("====== PART I ======")
	day02_I()
	fmt.Println("")
	fmt.Println("====== PART II ======")
	day02_II()
}
