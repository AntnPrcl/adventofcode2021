package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MAP_SIZE int

type point struct {
	X int
	Y int
}

type line struct {
	start point
	end   point
}

type matrix struct {
	Rows []row
}

type row struct {
	Column []int
}

func read_data() []line {
	input_byte, _ := os.ReadFile("inputs/day05.txt")
	input := strings.Split(string(input_byte), "\n")
	var data []line
	for _, row := range input {
		var new_row line
		fmt.Sscanf(row, "%d,%d -> %d,%d", &new_row.start.X, &new_row.start.Y, &new_row.end.X, &new_row.end.Y)
		data = append(data, new_row)
	}
	return data
}

func filter_data(old_input []line) []line {
	var new_data []line
	for _, line := range old_input {
		if (line.start.X == line.end.X) || (line.start.Y == line.end.Y) {
			new_data = append(new_data, line)
		}
	}
	return new_data
}

func generate_map() matrix {
	var new_map matrix
	for i := 0; i < MAP_SIZE; i++ {
		var new_row row
		for j := 0; j < MAP_SIZE; j++ {
			new_row.Column = append(new_row.Column, 0)
		}
		new_map.Rows = append(new_map.Rows, new_row)
	}
	return new_map
}
func count_cell(mapp matrix) int {
	count := 0
	for _, i := range mapp.Rows {
		for _, j := range i.Column {
			if j > 1 {
				count += 1
			}
		}
	}
	return count
}

func write_lines(input []line) {
	mapp := generate_map()
	for _, line := range input {
		if line.start.X == line.end.X {
			if line.start.Y < line.end.Y {
				for i := line.start.Y; i <= line.end.Y; i++ {
					mapp.Rows[i].Column[line.start.X] += 1
				}
			} else {
				for i := line.end.Y; i <= line.start.Y; i++ {
					mapp.Rows[i].Column[line.start.X] += 1
				}
			}
		} else {
			if line.start.X < line.end.X {
				for i := line.start.X; i <= line.end.X; i++ {
					mapp.Rows[line.start.Y].Column[i] += 1
				}
			} else {
				for i := line.end.X; i <= line.start.X; i++ {
					mapp.Rows[line.start.Y].Column[i] += 1
				}
			}
		}
	}
	for _, line := range mapp.Rows {
		fmt.Println(line)
	}
	fmt.Printf("Ans : " + strconv.Itoa(count_cell(mapp)))
}

func write_lines_II(input []line) {
	mapp := generate_map()
	for _, line := range input {
		if line.start.X == line.end.X || line.start.Y == line.end.Y {
			if line.start.X == line.end.X {
				if line.start.Y < line.end.Y {
					for i := line.start.Y; i <= line.end.Y; i++ {
						mapp.Rows[i].Column[line.start.X] += 1
					}
				} else {
					for i := line.end.Y; i <= line.start.Y; i++ {
						mapp.Rows[i].Column[line.start.X] += 1
					}
				}
			} else {
				if line.start.X < line.end.X {
					for i := line.start.X; i <= line.end.X; i++ {
						mapp.Rows[line.start.Y].Column[i] += 1
					}
				} else {
					for i := line.end.X; i <= line.start.X; i++ {
						mapp.Rows[line.start.Y].Column[i] += 1
					}
				}
			}
		} else {
			if line.start.X < line.end.X {
				line_size := line.end.X - line.start.X
				if line.start.Y < line.end.Y {
					for i := 0; i <= line_size; i++ {
						mapp.Rows[line.start.Y+i].Column[line.start.X+i] += 1
					}
				} else {
					for i := 0; i <= line_size; i++ {
						mapp.Rows[line.start.Y-i].Column[line.start.X+i] += 1
					}
				}
			} else {
				line_size := line.start.X - line.end.X
				if line.start.Y < line.end.Y {
					for i := 0; i <= line_size; i++ {
						mapp.Rows[line.start.Y+i].Column[line.start.X-i] += 1
					}
				} else {
					for i := 0; i <= line_size; i++ {
						mapp.Rows[line.start.Y-i].Column[line.start.X-i] += 1
					}
				}
			}
		}
	}

	fmt.Printf("Ans : " + strconv.Itoa(count_cell(mapp)))
}
func main() {
	MAP_SIZE = 10000
	data := read_data()
	//data = filter_data(data)
	write_lines_II(data)
}
