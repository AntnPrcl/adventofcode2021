package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func next_greater(x int, y int) bool {
	if x < y {
		return true
	} else {
		return false
	}
}

func read_input() (nums []int, err error) {
	b, err := os.ReadFile("../inputs/day01.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func day01_I() {
	test := [10]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	count := 0
	for i := 0; i < 9; i++ {
		if next_greater(test[i], test[i+1]) {
			count += 1
		}
	}
	input, _ := read_input()
	fmt.Println("Test Counter :", count)
	count = 0
	for i := 0; i < len(input)-1; i++ {
		if next_greater(input[i], input[i+1]) {
			count += 1
		}
	}
	fmt.Println("Input Counter :", count)
}

func day01_II() {
	test := [10]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	count := 0
	for i := 0; i < 9; i++ {
		if next_greater(test[i], test[i+1]) {
			count += 1
		}
	}
	input, _ := read_input()
	fmt.Println("Test Counter :", count)
	count = 0
	for i := 0; i < len(input)-3; i++ {
		a := input[i] + input[i+1] + input[i+2]
		b := input[i+1] + input[i+2] + input[i+3]
		if next_greater(a, b) {
			count += 1
		}
	}
	fmt.Println("Input Counter :", count)
}

func main() {
	fmt.Println("====== PART I ======")
	day01_I()
	fmt.Println("\n====== PART I ======")
	day01_II()
}
