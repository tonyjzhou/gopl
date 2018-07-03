// +build gcd

package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(gcd(6, 8))
	fmt.Println(gcd(545, 125))
}
