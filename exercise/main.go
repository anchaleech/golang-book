package main

import "fmt"

func main() {
	weatherCelsius(25, "Bangkok few cloud")
	weatherCelsius(34, "Tak sunny")
	weatherCelsius(17, "Phuket rainy")
	weatherCelsius(9, "Chaing-mai cold")
}

func weatherCelsius(c int, city string) {
	cnumber(c)
	fmt.Println(city)
}

func cnumber(c int) {
	switch c {
		case 9: print9()
		case 17: print17()
		case 25: print25()
		case 34: print34()
	}
}

func print9() {
	fmt.Println(" -")
	fmt.Println("| |")
	fmt.Println(" -")
	fmt.Println("  | c")
}

func print17() {
	fmt.Println("|  --")
	fmt.Println("|    | c")
}

func print25() {
	fmt.Println(" -   -")
	fmt.Println("  | |")
	fmt.Println(" -   -")
	fmt.Println("|     |")
	fmt.Println(" -   -  c")
}

func print34() {
	fmt.Println("-")
	fmt.Println("_| |_|")
	fmt.Println("_|   | c")
}