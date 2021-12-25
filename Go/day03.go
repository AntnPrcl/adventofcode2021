package day03

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_input() (input [][]int) {
	b, _ := os.ReadFile("../inputs/day03.txt")

	lines := strings.Split(string(b), "\n")

	//	input := make([]course, 0, len(lines))
	line := []int{}
	for _, l := range lines {

		if len(l) == 0 {
			continue
		}
		for i := 0; i < len(l); i++ {
			bit, _ := strconv.Atoi(l[i : i+1])
			line = append(line, bit)

		}

		input = append(input, line)
		line = []int{}
	}
	return input
}

func most_common_byte(input [][]int, index int) int {
	n1 := 0
	n0 := 0

	for i := 0; i < len(input); i++ {
		if input[i][index] == 0 {
			n0++
		} else {
			n1++
		}
	}
	if n1 >= n0 {
		return 1
	} else {
		return 0
	}
}

func less_common_byte(input [][]int, index int) int {
	n1 := 0
	n0 := 0
	for i := 0; i < len(input); i++ {
		if input[i][index] == 0 {
			n0++
		} else {
			n1++
		}
	}
	if n1 >= n0 {
		return 0
	} else {
		return 1
	}
}

func filter_most_common(input [][]int, index int) [][]int {
	bit := most_common_byte(input, index)
	var input_copy [][]int
	for i := 0; i < len(input); i++ {
		if input[i][index] == bit {
			input_copy = append(input_copy, input[i])
		}
	}
	return input_copy
}

func filter_less_common(input [][]int, index int) [][]int {
	bit := less_common_byte(input, index)
	var input_copy [][]int
	for i := 0; i < len(input); i++ {
		if input[i][index] == bit {
			input_copy = append(input_copy, input[i])
		}
	}
	return input_copy
}

func convert_result(input []int) int {
	var result string
	for i := 0; i < len(input); i++ {
		result = result + strconv.Itoa(input[i])
	}
	result_int, _ := strconv.ParseInt(result, 2, 64)
	return int(result_int)
}

func day03_I() {
	//test := [][]int{
	//	{0, 0, 1, 0, 0},
	//	{1, 1, 1, 1, 0},
	//	{1, 0, 1, 1, 0},
	//	{1, 0, 1, 1, 1},
	//	{1, 0, 1, 0, 1},
	//	{0, 1, 1, 1, 1},
	//	{0, 0, 1, 1, 1},
	//	{1, 1, 1, 0, 0},
	//	{1, 0, 0, 0, 0},
	//	{1, 1, 0, 0, 1},
	//	{0, 0, 0, 1, 0},
	//	{0, 1, 0, 1, 0}}
	input := read_input()
	gamma_rate := []int{}
	epsilon_rate := []int{}
	for i := 0; i < len(input[0]); i++ {
		gamma_rate = append(gamma_rate, most_common_byte(input, i))
		epsilon_rate = append(epsilon_rate, less_common_byte(input, i))
	}
	result := convert_result(gamma_rate) * convert_result(epsilon_rate)
	fmt.Println("Ans : ", result)
}

func day03_II() {

	input := read_input()
	input_copy := input
	for i := 0; i < len(input_copy[0]); i++ {
		input_copy = filter_most_common(input_copy, i)
		if len(input_copy) == 1 {
			break
		}
	}
	oxygen_generator_rating := convert_result(input_copy[0])
	fmt.Println("oxygen generator rating :", oxygen_generator_rating)
	input_copy = input
	for i := 0; i < len(input_copy[0]); i++ {
		input_copy = filter_less_common(input_copy, i)
		if len(input_copy) == 1 {
			break
		}
	}
	co2_scrubber_rating := convert_result(input_copy[0])
	fmt.Println("CO2 scrubber rating :", co2_scrubber_rating)
	fmt.Println("Answer :", co2_scrubber_rating*oxygen_generator_rating)
}
func main() {
	fmt.Println("*** ADVENT OF CODE 2021 ** DAY 03 ***")
	fmt.Println("====== PART I ======")
	day03_I()
	fmt.Println("====== PART II ======")
	day03_II()
}
