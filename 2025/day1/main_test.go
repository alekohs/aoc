package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	rows := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	expected := 3
	result := part1(rows)

	if result != expected {
		t.Fatalf("part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	rows := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	expected := 6
	result := part2(rows)

	if result != expected {
		t.Fatalf("part2() = %v, expected %v", result, expected)
	}

}
