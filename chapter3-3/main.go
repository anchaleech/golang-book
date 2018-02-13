package main

import "fmt"

func main() {
	fmt.Println("=====Floating point======")
	third := 1.0 / 3.0
	fmt.Println("third = %v\n", third)
	fmt.Println("third + third + third = %v\n", third+third+third)

	fmt.Println("====Comparing Floating Point=====")
	fmt.Println("0.1 + 0.2 == 0.3 is", 0.1+0.2 ==0.3)
	num := 0.1
	num += 0.2
	fmt.Println("num == 0.3 is", num == 0.3)
	fmt.Println("num is", num)
}