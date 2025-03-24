package main

import (
	// "errors"
	"fmt"
	// "regexp"
	"strconv"
	"strings"

	"github.com/alekohs/aoc/util"
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
	strData := strings.Split(data, "\n\n")
	rules := util.GetMatches(`\d+\|\d+`, strData[0])
	upgrades := strings.Split(util.GetMatches(`(?s)\d.*,\d\d`, strData[1])[0], "\n")

	result := 0

	for _, upgrade := range upgrades {
		arr := strings.Split(upgrade, ",")
		upgradeComplete := true

		for _, rule := range rules {
			ruleArray := strings.Split(rule, "|")
			intRules := toIntArray(ruleArray)

			fmt.Println(upgrade)
			fmt.Println(intRules)
			if !validateRule(intRules[0], intRules[1], arr) {
				upgradeComplete = false
				fmt.Println("Rule could not be validated")
				break
			}
		}

		if upgradeComplete {
		  fmt.Println("updated")
			result++
		}

	}

	fmt.Println(result)
	return result
}

func part2(data string) int {
	return 0
}

func validateRule(r1 int, r2 int, data []string) bool {
	valid := false
	row := toIntArray(data)
	i1 := -1
	i2 := -1

	for i, r := range row {
		if r1 == r {
			i1 = i
		}

		if r2 == r {
			i2 = i
		}
	}

	if ((i1 == -1 || i2 == -1) || i1 > i2) {
		valid = true
	}

	return valid
}

func toIntArray(data []string) []int {
	intArray := make([]int, len(data))

	for i, str := range data {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			fmt.Println("Error converrtint to int array")
		}

		intArray[i] = num
	}

	return intArray
}
