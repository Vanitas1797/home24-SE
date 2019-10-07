package main

import (
	"fmt"
)

func main() {
	s := "Das Auto"
	fmt.Println(s)
	s = reverse(s)
	fmt.Println(s)
}

func reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}