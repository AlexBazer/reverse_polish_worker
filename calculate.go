package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pop(a []float32) (float32, []float32, error) {
	l := len(a)
	if l == 0 {
		return 0, a, fmt.Errorf("Expression is not compnete and can't be parsed")
	}
	return a[l-1], a[:l-1], nil
}

func apply(s string, stack []float32) ([]float32, error) {
	var a, b, res float32
	var errA, errB error
	a, stack, errA = pop(stack)
	b, stack, errB = pop(stack)
	if errA != nil {
		return stack, errA
	} else if errB != nil {
		return stack, errB
	}

	switch {
	case s == "+":
		res = b + a
	case s == "-":
		res = b - a
	case s == "*":
		res = b * a
	case s == "/":
		res = b / a
	}
	stack = append(stack, res)
	return stack, nil
}

func calculate(expression string) (float32, error) {
	var splited = strings.Split(expression, " ")
	var stack = make([]float32, 0)
	var err error
	for _, s := range splited {
		switch s {
		case "+",
			"-",
			"*",
			"/":
			stack, err = apply(s, stack)
			if err != nil {
				return 0, err
			}
		default:
			digit, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return 0, err
			}
			stack = append(stack, float32(digit))
		}
	}
	if len(stack) == 1 {
		res, _, _ := pop(stack)
		return res, nil
	}
	return 0, fmt.Errorf("Expression is not complete and can't be parsed")
}
