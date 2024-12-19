package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func WelcomeMessage(day string) {
	fmt.Printf("Welcome to Advent of Code 2024 - %s! \n", day)
}

func ReadDataWithFunc[T any](path string, fn func(d string) T) T {
	if !strings.HasSuffix(path, "data.txt") {
		path += "/data.txt"
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return fn(string(data))
}

func ReadData(path string) []string {
	return ReadDataWithFunc(path, func(d string) []string {
		result := []string{}
		for _, r := range strings.Split(string(d), "\n") {
			if len(r) <= 0 {
				continue
			}

			result = append(result, r)
		}

		return result
	})
}

func ExtractIntArray(data []string) [][]int {
	result := [][]int{}

	for _, line := range data {
		lineArray := []int{}
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

func Contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}

	return false
}

func GetMatches(exp string, str string) []string {
	re, err := regexp.Compile(exp)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		panic(err)
	}

	matches := re.FindAllString(str, -1)
	if len(matches) > 0 {
		return matches
	}

	return []string{}
}

func ReplaceAll(exp string, str string, repl string) string {
	re, err := regexp.Compile(exp)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		panic(err)
	}

	return re.ReplaceAllString(str, repl)
}
