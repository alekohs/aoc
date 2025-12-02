package main

import (
	"fmt"
	"github.com/alekohs/aoc/util"
	"regexp"
	"strconv"
)

type aocPassword struct {
	ZeroPasses int
	ZeroEnd    int
}

func main() {
	util.WelcomeMessage("Day 1")
	data := util.ReadData("data.txt")

	fmt.Println("Part 1 ->  sum: ", part1(data))
	fmt.Println("Part 2 ->  sum: ", part2(data))

}

func part1(data []string) int {
	p := get_password(data)
	return p.ZeroEnd
}

func part2(data []string) int {
	p := get_password(data)
	return p.ZeroPasses
}

func get_password(data []string) aocPassword {
	current := 50
	re := regexp.MustCompile(`([A-Za-z]+)|([0-9]+)`)

	p := aocPassword{ZeroPasses: 0, ZeroEnd: 0}

	for _, r := range data {
		matches := re.FindAllStringSubmatch(r, -1)
		dir := ""
		num := 0
		for _, m := range matches {
			if m[1] != "" {
				dir = m[1]
			} else if m[2] != "" {
				l_num, numErr := strconv.Atoi(m[2])
				if numErr != nil {
					continue
				}
				num = l_num
			}
		}

		// fmt.Println("Current rotation", current, num)
		local_p, lc := rotate(current, dir, num)
		current = lc
		p.ZeroPasses += local_p.ZeroPasses
		p.ZeroEnd += local_p.ZeroEnd
	}

	return p
}

func rotate(current int, dir string, num int) (aocPassword, int) {
	min_num := 0
	max_num := 99
	multiplier := 1

	// fmt.Println("Pre calc", current, num, dir)

	if dir == "L" {
		multiplier = -1
	}

	zero_passes := 0
	for _ = range num {
		if current == 0 {
			zero_passes += 1
		}

		current += multiplier

		if current > max_num {
			current = 0
		} else if current < min_num {
			current = 99
		}
	}

	r := aocPassword{ZeroPasses: zero_passes }
	// fmt.Println("Calc complete", current)
	if (current == 0) {
		r.ZeroEnd = 1
	}
	return r, current
}
