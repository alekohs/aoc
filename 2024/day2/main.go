package main

import (
    "fmt"
    "math"
    "github.com/alekohs/aoc/util"
    "slices"
)
func main() {
    util.WelcomeMessage("Day 2");
    data := util.ReadData("data.txt");

    fmt.Println("Part 1 ->  sum: ", part1(data))
    fmt.Println("Part 2 ->  sum: ", part2(data))
}


func part1(data []string) int {
    lists := util.ExtractIntArray(data)
    result := 0

    for _, arr := range lists {
        if !isValid(arr) {
            continue
        }

        result++
    }

    return result
}

func part2(data []string) int {
    lists := util.ExtractIntArray(data)
    result := 0

    for _, arr := range lists {
        if isValid(arr) {
            result++
            continue
        }

        for i := 0; i < len(arr); i++ {
            tmpArr := make([]int, len(arr))
            copy(tmpArr[:], arr[:])
            arr3 := slices.Delete(tmpArr, i, i + 1)
            if !isValid(arr3) {
                continue
            }

            result++
            break;
        }
    }


    return result
}

func isValid(arr []int) bool {
    if (len(arr) <= 0) {
        return false
    }

    validNumbers := []int{ 1, 2, 3 }
    a := arr[0]
    increasing := arr[0] < arr[1]

    for i := 1; i < len(arr); i++ {
        b := arr[i]
        diff := math.Abs(float64(a - b))

        if !util.Contains(validNumbers, int(diff)) {
            return false
        }

        if (increasing && b < a) || (!increasing && b > a) {
            return false
        }

        a = b
    }

    return true
}
