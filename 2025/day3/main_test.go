package main

import (
	"testing"
)

var rows = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func TestPart1(t *testing.T) {
	expected := 357
	result := part1(rows)

	if result != expected {
		t.Fatalf("part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 3121910778619
	result := part2(rows)

	if result != expected {
		t.Fatalf("part2() = %v, expected %v", result, expected)
	}
}
