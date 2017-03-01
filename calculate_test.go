package main

import (
	"testing"
)

func TestReversePolish(t *testing.T) {
	var input = "5 1 2 + 4 * + 3 -"
	var output = 14.0
	var result = calculate(input)
	if output != result {
		t.Errorf("Get %v. But evaluation for %v should be %v", result, input, output)
	}
}
