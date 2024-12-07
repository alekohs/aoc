package main

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/alexanderkohkoinen/aoc/util"
	// "slices"
)

func main() {
	util.WelcomeMessage("Day 2")
	data := util.ReadData("data.txt")

	fmt.Println("Part 1 ->  sum: ", part1(data))
	fmt.Println("Part 2 ->  sum: ", part2(data))
}

func part1(lists []string) int {
	expr := `mul\((\d{1,3}\,\d{1,3}\))`
	sum := 0

	combined := strings.Join(lists, "\n")
	result := util.GetMatches(expr, combined)
	eval := BuildAlgoritm(result)
	sum += Evalute(eval)

	return sum
}

func part2(lists []string) int {
	expr := `mul\(\d{1,3}\,\d{1,3}\)|do\(\)|don't\(\)`
	sum := 0

	combined := strings.Join(lists, "\n")
	result := util.GetMatches(expr, combined)
	eval := BuildAlgoritm2(result)
	sum += Evalute(eval)

	return sum
}

func BuildAlgoritm(arr []string) string {
	result := ""

	for _, eval := range arr {
		str := util.ReplaceAll(`mul|\(|\)`, eval, "")
		// fmt.Println(str)

		str = util.ReplaceAll(`,`, str, `*`)
		if len(result) > 0 {
			result += " + "
		}

		result += str
	}

	return result
}

func BuildAlgoritm2(arr []string) string {
	result := ""
	skip := false

	for _, eval := range arr {
		str := util.ReplaceAll(`mul|\(|\)`, eval, "")
		switch str {
		case "don't":
			skip = true
		case "do":
			skip = false
			continue
		}

		// Skip current row
		if skip {
			continue
		}

		str = util.ReplaceAll(`,`, str, `*`)
		if len(result) > 0 {
			result += " + "
		}

		result += str
	}

	// fmt.Println(result)
	return result
}

func Evalute(value string) int {
	expr, err := govaluate.NewEvaluableExpression(value)
	if err != nil {
		fmt.Println("Error parsing: ", err)
		return 0
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		fmt.Println("Error evaluating: ", err)
		return 0
	}

	// fmt.Println("Result: ", result)

	num, ok := result.(float64)
	if !ok {
		return 0
	}

	return int(num)
}
