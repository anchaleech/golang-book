package main

import "fmt"

func main() {
	slice := []int{1,2,3}
	fmt.Println(slice)
	newSlice := make([]int, 2)
	fmt.Println(newSlice)
	copy(slice, newSlice)
	fmt.Println("slice: %v\n", slice)
	fmt.Println("slice: %v\n", newSlice)
}