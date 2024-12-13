package main

import (
	"errors"
	"fmt"
	"strings"

	// "strings"
	"github.com/alexanderkohkoinen/aoc/util"
)

func main() {
	util.WelcomeMessage("Day 4")
	data := util.ReadData("data.txt")

	fmt.Println("Part 1 ->  sum: ", part1(data))
	fmt.Println("Part 2 ->  sum: ", part2(data))
}

func part1(rows []string) int {
	grid := make([][]string, len(rows))
	word := []string{"X", "M", "A", "S"}
	sum := 0

	for i, row := range rows {
		grid[i] = splitOnChar(row)
	}

	directions := []Direction{
		East,
		West,
		North,
		South,
		SouthEast,
		SouthWest,
		NorthEast,
		NorthWest,
	}

	// Loop grid
	for r, row := range grid {
		indices := findOccurences(row, word[0])

		// fmt.Println(row, indices)
		for _, c := range indices {
			for _, dir := range directions {
				w1 := []string { "X" }
				cr := r
				cc := c

				validRows := [][]int{{cr, cc}}

				for j := 1; j < 4; j++ {
					v, rc, err := getValue(grid, cr, cc, dir)

					if err != nil {
                        continue
					}

					w1 = append(w1, v)
                    cr = rc[0]
					cc = rc[1]
                    validRows = append(validRows, []int{cr, cc})
				}

				if strings.Join(w1, "") == strings.Join(word, "") {
					// fmt.Println(validRows)
					sum += 1
				}
			}
		}
	}

	return sum
}

func part2(lists []string) int {
	sum := 0

	return sum
}

func splitOnChar(str string) []string {
	str = strings.TrimSpace(string(str))

	parts := []rune(str)
	result := make([]string, len(parts))
	for i, r := range parts {
		result[i] = string(r)
	}

	return result
}

func findOccurences(arr []string, letter string) []int {
	var indices []int
	for i, v := range arr {
		if v == letter {
			indices = append(indices, i)
		}
	}
	return indices
}

func getValue(grid [][]string, r int, c int, dir Direction) (string, []int, error) {
	maxRow := len(grid) - 1
	maxColumn := len(grid[0]) - 1

	switch dir {
	case North:
		r -= 1
		c += 0
	case NorthEast:
		r -= 1
		c += 1
	case NorthWest:
		r -= 1
		c -= 1
	case East:
		r += 0
		c += 1
	case South:
		r += 1
		c += 0
	case SouthEast:
		r += 1
		c += 1
	case SouthWest:
		r += 1
		c -= 1
	case West:
		r += 0
		c += -1

	}

	// Index of out bounds
	if r < 0 || r > maxRow {
        message := fmt.Sprintf("row out of bounds %d > %d", r, maxRow)
		return "", []int{ r, c }, errors.New(message)
	}

    if c < 0 || c > maxColumn  {
        message := fmt.Sprintf("column out of bounds %d > %d", c, maxColumn)
		return "", []int{ r, c }, errors.New(message)
    }

	return grid[r][c], []int{r, c}, nil
}

type Direction int

const (
	North Direction = iota
	NorthEast
	NorthWest
	East
	South
	SouthEast
	SouthWest
	West
)
