package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	rows := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
		"565653-565659",
		"824824821-824824827",
		"2121212118-212121212",
	}

	expected := int64(1227775554)
	result := part1(rows)

	fmt.Println(result)
	if result != expected {
		t.Fatalf("part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	rows := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
		"565653-565659",
		"824824821-824824827",
		"2121212118-2121212124",
	}

	expected := int64(4174379265)
	result := part2(rows)

	if result != expected {
		t.Fatalf("part2() = %v, expected %v", result, expected)
	}

}
