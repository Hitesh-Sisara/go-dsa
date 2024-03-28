package main

import (
	"fmt"
	"time"
)

func main() {
	n := 40

	start := time.Now()
	result := fibonaccir(n)
	timetaken := time.Since(start)
	fmt.Println("Recursive fibonacci(", n, ") result:", result, "Time:", timetaken)

	start = time.Now()
	result = fibonacci(n)
	timetaken = time.Since(start)
	fmt.Println(" fibonacci(", n, ") result:", result, "Time:", timetaken)
}

func fibonaccir(n int) int {
	if n <= 1 {
		return 1
	}

	return fibonaccir(n-1) + fibonaccir(n-2)
}

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}

	a, b := 0, 1

	for i := 2; i <= n; i++ {

		c := a + b
		a = b
		b = c

	}

	return b
}
