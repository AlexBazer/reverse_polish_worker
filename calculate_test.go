package main

import (
	"testing"
)

func TestReversePolish(t *testing.T) {
	checkCalculation(t, "5 1 +", 6)
	checkCalculation(t, "5.2 1.4 +", 6.6)
	checkCalculation(t, "2 21 * 30 -", 12)
	checkCalculation(t, "5 1 2 + 4 * + 3 -", 14)
}

func TestReversePolishErrors(t *testing.T) {
	checkCalculationOnError(t, "5 1 +-")
	checkCalculationOnError(t, "5 1 + -")
	checkCalculationOnError(t, "5 +")
}

func checkCalculation(t *testing.T, input string, output float32) {
	result, err := calculate(input)
	if err != nil {
		t.Error(err)
	} else if float32(output) != float32(result) {
		t.Errorf("Get %v. But evaluation for %v should be %v", result, input, output)
	}
}
func checkCalculationOnError(t *testing.T, input string) {
	_, err := calculate(input)
	if err == nil {
		t.Errorf("Errors should be return for this input: %v", input)
	}
}
