package main

import "fmt"

func main() {

	printToConsole()
}

func printToConsole() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
