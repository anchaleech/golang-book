package main

import "fmt"

func main() {
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println("%v\n", string(c))
	}
}