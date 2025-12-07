package main

import (
	"fmt"
	"strings"

	"github.com/alekohs/aoc/util"
)

func main() {
	util.WelcomeMessage("Day 4")
	data := util.ReadData("data.txt")

	fmt.Println("Part 1 ->  sum: ", part1(data))
	fmt.Println("Part 2 ->  sum: ", part2(data))

}

func part1(data []string) int {
	num := 0
	matrix := make([][]string, 0, len(data))
	row_length := len(data)
	col_length := len(data[0])

	for _, line := range data {
		line = strings.TrimSpace(line)
		arr := make([]string, 0, len(line))
		for _, ch := range line {
			arr = append(arr, string(ch))
		}
		matrix = append(matrix, arr)
	}

	for r := range row_length {
		for c := range col_length {
			if !validate(matrix, r, c) {
				continue
			}

			occurrences := 0
			dirs := [][2]int{
				{1, 0}, {1, 1}, {1, -1},
				{-1, 0}, {-1, 1}, {-1, -1},
				{0, 1}, {0, -1},
			}

			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if validate(matrix, nr, nc) {
					occurrences++
				}
			}

			if occurrences < 4 {
				num++
			}
		}
	}

	return num
}

func part2(data []string) int {
	num := 0
	matrix := make([][]string, 0, len(data))
	row_length := len(data)
	col_length := len(data[0])

	for _, line := range data {
		line = strings.TrimSpace(line)
		arr := make([]string, 0, len(line))
		for _, ch := range line {
			arr = append(arr, string(ch))
		}
		matrix = append(matrix, arr)
	}

	loop := 0
	for loop < 1 {
		match := false

		for r := range row_length {
			for c := range col_length {
				if !validate(matrix, r, c) {
					continue
				}

				occurrences := 0
				dirs := [][2]int{
					{1, 0}, {1, 1}, {1, -1},
					{-1, 0}, {-1, 1}, {-1, -1},
					{0, 1}, {0, -1},
				}

				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if validate(matrix, nr, nc) {
						occurrences++
					}
				}

				if occurrences < 4 {
					num++
					matrix[r][c] = "."
					match = true
				}
			}
		}

		if !match {
			loop = 1
		}
	}

	return num
}

func validate(matrix [][]string, row int, col int) bool {
	row_max := len(matrix) - 1
	col_max := len(matrix[0]) - 1

	if row < 0 || col < 0 || row > row_max || col > col_max {
		return false
	}
	if matrix[row][col] == "@" {
		return true
	}
	return false
}
