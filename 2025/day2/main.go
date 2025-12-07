package main

import (
	"fmt"
	"github.com/alekohs/aoc/util"
	"strconv"
	"strings"
)

func main() {
	util.WelcomeMessage("Day 1")
	data := util.ReadData("data.txt")
	ids := strings.Split(data[0], ",")

	fmt.Println("Part 1 ->  sum: ", part1(ids))
	fmt.Println("Part 2 ->  sum: ", part2(ids))

}

func part1(data []string) int64 {
	num := get_duplicate_digits(data, find_recurring_sequence)
	return num
}

func part2(data []string) int64 {
	num := get_duplicate_digits(data, find_recurring_sequence_repeat)
	return num
}

func get_duplicate_digits(data []string, re func(string) int64) int64 {
	var num int64
	for _, line := range data {
		digits := strings.Split(line, "-")
		from, fromErr := strconv.ParseInt(digits[0], 10, 64)
		to, toErr := strconv.ParseInt(digits[1], 10, 64)
		if fromErr != nil || toErr != nil {
			return 0
		}

		for i := from; i <= to; i++ {
			str := strconv.FormatInt(i, 10)
			num += re(str)
		}
	}

	return num
}

func find_recurring_sequence(digits string) int64 {
	n := len(digits)
	if n < 2 || n%2 != 0 {
		return 0
	}

	half := n / 2
	if digits[:half] != digits[half:] {
		return 0
	}

	r, _ := strconv.ParseInt(digits, 10, 64)
	return r
}

func find_recurring_sequence_repeat(digits string) int64 {
	n := len(digits)
	if n < 2 {
		return 0
	}

	half := n / 2
	for i := 1; i <= half; i++ {
		if n%i != 0 {
			continue
		}

		chunk := digits[:i]
		if strings.Repeat(chunk, n/i) == digits {
			r, _ := strconv.ParseInt(digits, 10, 64)
			return r
		}
	}

	return 0
}
