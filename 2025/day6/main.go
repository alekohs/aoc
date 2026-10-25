package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/alekohs/aoc/util"
)

type interval struct{ start, end int }

func main() {
	util.WelcomeMessage("Day 6")
	data := util.ReadData("data.txt")

	fmt.Println("Part 1 -> sum: ", part1(data))
	fmt.Println("Part 2 -> sum: ", part2(data))

}

func part1(data []string) int {
	rows := len(data)
	operatorIndex := rows - 1
	numIndex := rows - 2
	numbers := util.ExtractIntArray(data)
	operators := strings.ReplaceAll(data[operatorIndex], " ", "")
	columns := len(numbers[0])

	return getTotal(columns, numIndex, numbers, operators)
}

func part2(data []string) int {
	rows := len(data)
	operatorIndex := rows - 1
	numbers := getNumbers(data)
	slices.Reverse(numbers)

	operators := strings.ReplaceAll(data[operatorIndex], " ", "")
	columns := len(numbers)

	total := 0
	for c, oc := 0, columns-1; c < columns; c, oc = c+1, oc-1 {
		value := 0
		for i, n := range numbers[c] {
			if i == 0 {
				value = n
				continue
			}

			if operators[oc] == '*' {
				value = value * n
			} else if operators[oc] == '+' {
				value = value + n
			}
		}

		total += value
	}

	return total
}

func getNumbers(data []string) [][]int {
	result := [][]int{}
	numRows := data[:len(data)-1]

	width := 0
	for _, line := range numRows {
		width = max(width, len(line))
	}

	problem := []int{}
	for x := 0; x < width; x++ {
		num := 0
		hasDigit := false
		for _, line := range numRows {
			if x < len(line) && line[x] != ' ' {
				num = num*10 + int(line[x]-'0')
				hasDigit = true
			}
		}

		if hasDigit {
			problem = append(problem, num)
		} else {
			result = append(result, problem)
			problem = []int{}
		}
	}
	result = append(result, problem)

	return result
}

func getTotal(columns int, numIndex int, numbers [][]int, operators string) int {
	total := 0
	for c := 0; c < columns; c++ {
		value := 0

		for v := 0; v <= numIndex; v++ {
			if v == 0 {
				value = numbers[v][c]
				continue
			}

			if operators[c] == '*' {
				value = value * numbers[v][c]
			} else if operators[c] == '+' {
				value = value + numbers[v][c]
			}
		}

		total += value
	}

	return total
}
