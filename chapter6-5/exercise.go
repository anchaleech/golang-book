package main

import (
	"strconv"
	"fmt"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(number, fizzbuzz(number))
	}
}

func fizzbuzz(number int) string {
	ln := [3]int{15, 3, 5}
	str := [3]string{"FIzzbuzz", "Fizz", "Buzz"}
	for i := 0; i < len(ln); i++ {
		if number%ln[i] == 0 {
			return str[i]
		}
	}
	return strconv.Itoa(number)
}