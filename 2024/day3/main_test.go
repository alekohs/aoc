
package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	rows := []string {
        "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}

	expected := 2
	result := part1(rows)

	if result != expected {
		t.Fatalf("part1() = %v, want %v", result, expected)
	}
}


func TestPart2(t *testing.T) {
	rows := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	expected := 4
	result := part2(rows)

    if result != expected {
		t.Fatalf("part2() = %v, expected %v", result, expected)
	}

}
