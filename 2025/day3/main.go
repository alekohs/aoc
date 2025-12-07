package main

import (
	"fmt"
	"github.com/alekohs/aoc/util"
	"strconv"
)

func main() {
	util.WelcomeMessage("Day 1")
	data := util.ReadData("data.txt")

	fmt.Println("Part 1 ->  sum: ", part1(data))
	fmt.Println("Part 2 ->  sum: ", part2(data))

}

func part1(data []string) int {
	return get_banks(data, 2)
}

func part2(data []string) int {
	return get_banks(data, 12)
}

func get_banks(data []string, length int) int {
	num := 0
	for _, line := range data {
		num += get_largest(line, length)
	}
	return num
}

func get_largest(line string, keep int) int {
	remove := len(line) - keep
	stack := make([]rune, 0, len(line))

	for _, digit := range line {
		for len(stack) > 0 && stack[len(stack)-1] < digit && remove > 0 {
			stack = stack[:len(stack)-1]
			remove--
		}

		stack = append(stack, digit)
	}

	result := string(stack[:keep]);
	dt, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	return dt
}

func get_duplicates(line string) int {
	first := 0
	second := 0

	for i, digit := range line {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			continue
		}

		if i == len(line)-1 {
			if d > second {
				second = d
				continue
			}
		}

		if d > first {
			first = d
			second = 0
		} else if d > second {
			second = d
		}
	}

	total := fmt.Sprint(first) + fmt.Sprint(second)
	dt, err := strconv.Atoi(total)
	if err != nil {
		return 0
	}
	return dt
}
