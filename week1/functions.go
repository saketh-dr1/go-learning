package main

import "fmt"

func add(x, y int) (int, string) {
	return x + y, "hello"
}

func divide(num, deno float64) (float64, error) {
	if deno == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return num / deno, nil
}

// func as a parameter
func callFunc(callable func(int) int) int {
	return callable(10)
}

func doubleNummber(num int) int {
	return 2 * num
}

func getFunc(str string) func(string) string {
	return func(str2 string) string {
		return str + str2
	}
}

// Variadic func -- n no.of inputs
func sum(nums ...int) (s int) {

	for _, value := range nums {
		s += value
	}

	return
}
