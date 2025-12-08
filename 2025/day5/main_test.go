package main

import (
	"testing"
)

var rows = []string{
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

func TestPart1(t *testing.T) {
	expected := 3
	result := part1(rows)

	if result != expected {
		t.Fatalf("part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 14
	result := part2(rows)

	if result != expected {
		t.Fatalf("part2() = %v, expected %v", result, expected)
	}
}
