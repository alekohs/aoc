package main

import (
    "fmt"
    "github.com/alekohs/aoc/util"
    "strconv"
    "strings"
    "sort"
    "math"
)

func main() {
    util.WelcomeMessage("Day 1");
    var data = util.ReadData("day1/data.txt");

    Part1(data)
    Part2(data)
}


func Part1(data []string) {
    lists := ExtractData(data)
    sum := SortAndSum(lists[0], lists[1])
    fmt.Println("Part 1 ->  Sum is: ", sum)
}

func Part2(data []string) {
    lists := ExtractData(data)
    sum := 0

    for _, num := range lists[0] {
        counter := 0
        for _, num2 := range lists[1] {
            if num != num2 {
                continue
            }

            counter++
        }

        sum += num * counter
    }

    fmt.Println("Part 2 ->  Sum is: ", sum)
}

func ExtractData(data []string) [][]int {
    leftList := []int{} 
    rightList := []int{}

    for _, line := range data {
        parts := strings.Split(line, " ")
        if len(parts) < 2 {
            continue
        }

        // fmt.Println("", parts)
        num1, err1 := strconv.Atoi(parts[0])
        num2, err2 := strconv.Atoi(parts[3])

        // fmt.Printf("%s %s \n", err1, err2)

        if err1 == nil && err2 == nil {
            leftList = append(leftList, num1)
            rightList = append(rightList, num2)
        }
    }

    return [][]int {leftList, rightList}
}

func SortAndSum(a []int, b []int) int {
    // fmt.Println("Sort and sum the lists")

    sort.Ints(a)
    sort.Ints(b)

    sum := 0

    for i := 0; i < len(a); i++ {
        aNum := a[i]
        bNum := b[i]

        diff := math.Abs(float64(aNum - bNum))
        sum += int(diff)
    } 

    return sum
}
