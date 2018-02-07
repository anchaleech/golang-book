package main

import "fmt"

func main() {
	fmt.Println("=====String=====")
	backticks := `hello world, 
today's good day.`
	fmt.Println(backticks)

	doubleQote := "hello world, today's good day."
	fmt.Println(doubleQote)
}