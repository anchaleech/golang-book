package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		if number%15 == 0 {
			fmt.Println(number, "FizzBuzz")
		} else if number%3 ==0 {
			fmt.Println(number, "Fizz")
		} else if number%5 == 0 {
			fmt.Println(number, "Buzz")
		} else {
			fmt.Println(number)
		}
	}

	fmt.Println("======================")

	for number := 1; number <= 100; number++ {
		i := number%5
		switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			default: fmt.Println("Unknown Number")
			}
		}

	fmt.Println("======================")

	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 0: fmt.Println("even")
		default: fmt.Println("Unknown Number")
		}
	}
	}