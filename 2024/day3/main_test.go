
package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	rows := []string {
        "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}

	expected := 161
	result := part1(rows)

	if result != expected {
		t.Fatalf("part1() = %v, want %v", result, expected)
	}
}


func TestPart2(t *testing.T) {
	rows := []string{
        "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}

	expected := 48
	result := part2(rows)

    if result != expected {
		t.Fatalf("part2() = %v, expected %v", result, expected)
	}

}
