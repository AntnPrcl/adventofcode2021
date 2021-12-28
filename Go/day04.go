package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type cell struct {
	number int
	draw   bool
}

func remove(slice [][][]cell, s [][]cell) [][][]cell {
	for i := 0; i < len(slice); i++ {
		if reflect.DeepEqual(slice[i], s) {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func read_input() ([]int, [][][]cell) {
	input, _ := os.ReadFile("inputs/day04.txt")
	lines := strings.Split(string(input), "\n")
	nums := strings.Split(lines[0], ",")
	var new_nums []int
	for _, num := range nums {
		new_value, _ := strconv.Atoi(num)
		new_nums = append(new_nums, new_value)
	}
	tables := strings.Split(string(input), "\n\n")
	tables = tables[1:]
	var new_tables [][][]cell
	for _, table := range tables {
		table_slice := strings.Split(string(table), "\n")
		var new_table [][]cell
		for _, row := range table_slice {
			row = strings.Replace(row, "  ", " ", -1)
			row_slice := strings.Split(row, " ")
			if row_slice[0] == "" {
				row_slice = row_slice[1:]
			}
			var new_row []cell
			for _, column := range row_slice {
				var new_column cell
				new_column.draw = false
				new_column.number, _ = strconv.Atoi(column)
				new_row = append(new_row, new_column)
			}
			new_table = append(new_table, new_row)
		}
		new_tables = append(new_tables, new_table)
	}

	return new_nums, new_tables
}

func day04_I(tables [][][]cell, nums []int) {
	for _, num := range nums {
		for i, table := range tables {
			for j, row := range table {
				for k, column := range row {
					if column.number == num {
						tables[i][j][k].draw = true
					}
				}
			}
		}

		for _, table := range tables {
			for i := 0; i < 5; i++ {
				if (table[i][0].draw == true && table[i][1].draw == true && table[i][2].draw == true && table[i][3].draw == true && table[i][4].draw == true) || (table[0][i].draw == true && table[1][i].draw == true && table[2][i].draw == true && table[3][i].draw == true && table[4][i].draw == true) {
					score := 0
					for _, row := range table {
						for _, column := range row {
							if column.draw == false {
								score += column.number
							}
						}
					}
					score = score * num
					fmt.Println("Ans :", score)
					return
				}
			}
		}
	}
}

func day04_II(tables [][][]cell, nums []int) {
	for _, num := range nums {
		for i, table := range tables {
			for j, row := range table {
				for k, column := range row {
					if column.number == num {
						tables[i][j][k].draw = true
					}
				}
			}
		}

		if len(tables) != 1 {
			for _, table := range tables {
				for i := 0; i < 5; i++ {
					if (table[i][0].draw == true && table[i][1].draw == true && table[i][2].draw == true && table[i][3].draw == true && table[i][4].draw == true) || (table[0][i].draw == true && table[1][i].draw == true && table[2][i].draw == true && table[3][i].draw == true && table[4][i].draw == true) {

						tables = remove(tables, table)
						break
					}
				}
			}
		} else {
			table := tables[0]
			for i := 0; i < 5; i++ {
				if (table[i][0].draw == true && table[i][1].draw == true && table[i][2].draw == true && table[i][3].draw == true && table[i][4].draw == true) || (table[0][i].draw == true && table[1][i].draw == true && table[2][i].draw == true && table[3][i].draw == true && table[4][i].draw == true) {
					score := 0
					for _, row := range tables[0] {
						for _, column := range row {
							if column.draw == false {
								score += column.number
							}
						}
					}
					score = score * num
					fmt.Println(table)
					fmt.Println("Ans :", score)
					return
				}
			}
		}
	}
}

func main() {
	read_input()
	nums, tables := read_input()

	day04_II(tables, nums)
}
