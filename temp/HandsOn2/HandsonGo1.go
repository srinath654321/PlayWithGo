package main

import "fmt"

const (
	YEAR2016 = 2016 + iota
	YEAR2017 = 2016 + iota
	YEAR2018 = 2016 + iota
	YEAR2019 = 2016 + iota
)

func main() {

	printDecBinHexNumber()
	fmt.Println(YEAR2016, YEAR2017, YEAR2018, YEAR2019)
}

func printDecBinHexNumber() {

	X := 42

	fmt.Printf("%d\t%b\t%#x\n", X, X, X)

	Y := X << 1

	fmt.Printf("%d\t%b\t%#x\n", Y, Y, Y)

}
