// +build fib

package main

import "fmt"

func fib(n int) int {
	a, b := 1, 1

	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

func main() {
	for i := 1; i <= 6; i++ {
		fmt.Println(fib(i))
	}
}
