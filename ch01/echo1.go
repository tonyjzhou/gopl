// Print comamnd line arguemnts
// +build echo1

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, separator string
	for i := 1; i < len(os.Args); i++ {
		s += separator + os.Args[i]
		separator = " "
	}
	fmt.Println(s)
}
