package util

import (
    "fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)


func WelcomeMessage(day string) {
    fmt.Printf("Welcome to Advent of Code 2024 - %s! \n", day)
}


func ReadData(path string) []string {
    file, err := os.Open(path)
    if err != nil {
        log.Fatalf("unable to open file: %v", err)
    }

    data, err := io.ReadAll(file)
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

    return strings.Split(string(data), "\n")
}


func ExtractIntArray(data []string) [][]int {
    result := [][]int{}

    for _, line := range data {
        lineArray := []int {}
        parts := strings.Split(line, " ")
        for _, part := range parts {
            num1, err1 := strconv.Atoi(part)
            if err1 == nil {
                lineArray = append(lineArray, num1)
            }
        }
        result = append(result, lineArray)
    }

    return result
}
