package main

import (
	"fmt"
	"strings"

	// "strings"
	"github.com/alexanderkohkoinen/aoc/util"
)

func main() {
	util.WelcomeMessage("Day 2")
	data := util.ReadData("data.txt")

	fmt.Println("Part 1 ->  sum: ", part1(data))
	fmt.Println("Part 2 ->  sum: ", part2(data))
}

func part1(lists []string) int {
	sum := 0
	md := make([][]string, len(lists))

	for i, list := range lists {
		md[i] = splitOnChar(list)
	}

	for i, row := range md {
		first := false
		last := false

		if i == 0 { first := true }
        if i == len(md[i]) { last := true }

		for j := 0; j < len(row); j++ {
			fmt.Println("", j, i)
		}
	}

	fmt.Println(md)

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

