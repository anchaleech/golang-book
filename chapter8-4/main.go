package main

import "fmt"

func main() {
	array := [3]int{1,2,3}
	fmt.Printf("origin addr %p\n", &array)
	double(&array)
	fmt.Printf("original %v\n", array)
}

func double(nums *[3]int) {
	fmt.Printf("double addr %p\n", nums)
	fmt.Printf("double value %p\n", *nums)
	for i := range nums {
		//nums[i] *= 2
		(*nums)[i] *= 2
	}
	fmt.Println(nums)
}