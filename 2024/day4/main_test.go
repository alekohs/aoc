package main

import (
	"testing"
	"github.com/alexanderkohkoinen/aoc/util"
)

func TestPart1(t *testing.T) {
	data := util.ReadData("test-data.txt")

	expected := 18
	result := part1(data)

	if result != expected {
		t.Fatalf("part1() = %v, want %v", result, expected)
	}
}


func TestPart2(t *testing.T) {
	data := util.ReadData("test-data.txt")

	expected := 9
	result := part2(data)

	if result != expected {
		t.Fatalf("part2() = %v, want %v", result, expected)
	}

}
