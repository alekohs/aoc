package main

import (
	"testing"
)

var rows = []string{
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

func TestPart1(t *testing.T) {
	expected := 13
	result := part1(rows)

	if result != expected {
		t.Fatalf("part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 43
	result := part2(rows)

	if result != expected {
		t.Fatalf("part2() = %v, expected %v", result, expected)
	}
}
