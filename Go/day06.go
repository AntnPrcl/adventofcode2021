package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pass_on_day(aquarium [9]int) [9]int {
	new_aquarium := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 8; i++ {
		new_aquarium[i] += aquarium[i+1]
	}
	new_aquarium[8] += aquarium[0]
	new_aquarium[6] += aquarium[0]
	return new_aquarium
}

func read_input() [9]int {
	aquarium := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	input_string, _ := os.ReadFile("inputs/day06.txt")
	input := strings.Split(string(input_string), ",")
	for _, num := range input {
		index, _ := strconv.Atoi(num)
		aquarium[index] += 1
	}
	return aquarium
}

func count_fish(aquarium [9]int) int {
	count := 0
	for _, nb := range aquarium {
		count += nb
	}
	return count
}

func main() {
	aquarium := read_input()
	for i := 0; i < 256; i++ {
		aquarium = pass_on_day(aquarium)
	}
	fmt.Printf("Ans : %d", count_fish(aquarium))
}
