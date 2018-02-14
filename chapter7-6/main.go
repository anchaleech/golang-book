package main

import "fmt"

func main() {
	x := map[string]int{
		"one":	1,
		"two":	2,
		"three":	3,
	}
	fmt.Println(x)

	delete(x,"two")
	fmt.Println("After delete: %v\n", x)
}