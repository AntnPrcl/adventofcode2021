package main

import "fmt"

func next_greater(x int, y int) bool {
	if x < y {
		return true
	} else {
		return false
	}
}

func day01_I() {
	test := [10]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	count := 0
	for i := 0; i < 9; i++ {
		if next_greater(test[i], test[i+1]) {
			count += 1
		}
	}
	fmt.Println("Counter :", count)
}

func main() {
	day01_I()
}
