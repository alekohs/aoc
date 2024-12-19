package main

import (
	// "errors"
	"fmt"
	// "regexp"
	// "strings"

	// "strings"
	"github.com/alexanderkohkoinen/aoc/util"
)

func main() {
	util.WelcomeMessage("Day 5")
	data := util.ReadDataWithFunc[string]("data.txt", func(d string) string { 
        return d
    })

	fmt.Println("Part 1 ->  sum: ", part1(data))
	fmt.Println("Part 2 ->  sum: ", part2(data))
}

func part1(data string) int {
    rules := util.GetMatches(`\d+\|\d+`, data)
    pageNumbers := util.GetMatches(`^\d.*,\d$`, data)

    fmt.Println(rules, pageNumbers)
    return 0
}

func part2(data string) int {
    return 0
}

