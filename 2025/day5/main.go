package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/alekohs/aoc/util"
)

type interval struct{ start, end int }

func main() {
	util.WelcomeMessage("Day 5")
	data := util.ReadDataWithEmpty("data.txt")

	fmt.Println("Part 1 ->  sum: ", part1(data))
	fmt.Println("Part 2 ->  sum: ", part2(data))

}

func part1(data []string) int {
	num := 0
	groups := split_on_empty(data)
	ranges := get_ranges(groups[0])

	for _, g := range groups[1] {
		for _, r := range ranges {
			current, _ := strconv.Atoi(g)
			min := r[0]
			max := r[1]

			// Found a freshie
			if current >= min && current <= max {
				num++
				break
			}
		}
	}

	return num
}

func part2(data []string) int {
	groups := split_on_empty(data)
	ranges := get_ranges(groups[0])
	var intervals []interval

	for _, r := range ranges {
		intervals = append(intervals, interval{r[0], r[1]})
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	merged := []interval{}
	for _, curr := range intervals {
		if len(merged) == 0 {
			merged = append(merged, curr)
			continue
		}
		last := &merged[len(merged)-1]
		if curr.start <= last.end {
			if curr.end > last.end {
				last.end = curr.end
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return numbers_in_range(merged)
}

func get_ranges(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		digits := strings.Split(line, "-")
		var local_arr []int
		for _, d := range digits {
			value, err := strconv.Atoi(d)
			if err != nil {
				continue
			}
			local_arr = append(local_arr, value)
		}

		result = append(result, local_arr)
	}

	return result
}

func split_on_empty(lines []string) [][]string {
	var result [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				result = append(result, current)
				current = []string{}
			}
			continue
		}

		current = append(current, line)
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result
}

func numbers_in_range(merged []interval) int {
	counter := 0
	for _, r := range merged {
		counter += r.end - r.start + 1
	}

	return counter
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
