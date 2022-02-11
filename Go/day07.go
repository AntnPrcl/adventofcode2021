package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func read_input() []int {
	input_byte, _ := os.ReadFile("inputs/day07.txt")
	input_str := strings.Split(string(input_byte), ",")
	var input []int
	for _, nb := range input_str {
		tmp, _ := strconv.Atoi(nb)
		input = append(input, tmp)
	}
	return input
}

func count_fuel_cost(crabs []int, position int) int {
	count := float64(0)
	for _, crab := range crabs {
		count += math.Abs(float64(crab - position))
	}
	return int(count)
}

func count_fuel_cost_2(crabs []int, position int) int {
	count := 0
	for _, crab := range crabs {
		for i := 0; i <= int(math.Abs(float64(crab-position))); i++ {
			count += i
		}
	}
	return count
}

func main() {
	crabs := read_input()
	cheap_cost := count_fuel_cost(crabs, 0)
	for i := 1; i < 1904; i++ {
		fuel_cost := count_fuel_cost(crabs, i)
		if fuel_cost < cheap_cost {
			cheap_cost = fuel_cost
		}
	}
	fmt.Printf("Part I\n")
	fmt.Printf("Ans : %d\n\n", cheap_cost)

	cheap_cost = count_fuel_cost_2(crabs, 0)
	for i := 1; i < 1904; i++ {
		fuel_cost := count_fuel_cost_2(crabs, i)
		if fuel_cost < cheap_cost {
			cheap_cost = fuel_cost
		}
	}
	print("Part II\n")
	print("Ans : ", cheap_cost)
}
