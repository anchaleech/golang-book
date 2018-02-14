package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		if m15(number) == 0 {
			fmt.Println(number, "FizzBuzz")
		} else if m3(number) == 0 {
			fmt.Println(number, "Fizz")
		} else if m5(number) == 0 {
			fmt.Println(number, "Buzz")
		} else {
			fmt.Println(number)
		}
	}

	fmt.Println("====================")
	for number := 1; number <= 100; number++ {
		fmt.Println(number, modx(number))
	}

}

func modx(n int) string {
	var val string
	if (n % 15) == 0 {
		val = "FizzBuzz"
	} else if (n % 3) == 0 {
		val = "Fizz"
	} else if (n % 5) == 0 {
		val = "Buzz"
	}
	return val
}

func m15(n int) int {
	return n % 15
}

func m5(n int) int {
	return n % 5
}

func m3(n int) int {
	return n % 3
}
