package main

import (
    "fmt"
    "github.com/alexanderkohkoinen/aoc/util"
    "github.com/Knetic/govaluate"
    // "slices"
)

func main() {
    util.WelcomeMessage("Day 2");
    data := util.ReadData("data.txt");

    fmt.Println("Part 1 ->  sum: ", part1(data))
    fmt.Println("Part 2 ->  sum: ", part2(data))
}

func part1(lists []string) int {
    // expr := `(mul\((\d{2}\,\d{1}|\d{1}\,\d{1}|\d{1}\,\d{2})\))`
    expr := `mul\((\d{1,3}\,\d{1,3}\))`
    sum := 0

    for _, row := range lists {
        result := util.GetMatches(expr, row)
        fmt.Println(result)
        eval := BuildAlgoritm(result)
        fmt.Println(eval)
        sum += Evalute(eval)
    }

    return sum
}


func part2(arr []string) int {
    return 2
}


func BuildAlgoritm(arr []string) string {
    result := ""

    for _, eval := range arr {
        str := util.ReplaceAll(`mul|\(|\)`, eval, "")
        // fmt.Println(str)

        str = util.ReplaceAll(`,`,str, `*` )
        if len(result) > 0 {
            result += " + "
        }

        result += str
    }

    return result
}

func Evalute(value string) int {
    expr, err := govaluate.NewEvaluableExpression(value)
    if err != nil {
        fmt.Println("Error parsing", err)
        return 0
    }

    result, err := expr.Evaluate(nil)
    if err != nil {
        fmt.Println("Error evaluating", err)
        return 0
    }

    fmt.Println("Result: ", result)

    num, ok := result.(float64) 
    if !ok {
        return 0
    }

    return int(num)
}
