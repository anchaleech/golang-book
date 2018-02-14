package main

import "fmt"

func main() {
	fmt.Println(add(1,2,3))

	xs := []int{1,2,3,4}
	fmt.Println(add(xs...))
}

func add(arge ...int) int {
	total := 0
	for _, v := range arge {
		total += v
	}
	return total
}