package main

import (
	"fmt"
	"runtime"
)

var h int = 42
var i string = "James Bond"
var j bool = true
var O int

func main() {

	knowAboutVaribleScope()
	useType()
	convertStringToBytes()
}

func knowAboutVaribleScope() {

	s := fmt.Sprintf("%v\t%v\t%v", h, i, j)
	fmt.Println(s)

}

func useType() {

	type GAMESCORE int
	var N GAMESCORE
	fmt.Println(N)
	fmt.Printf("%T\n", N)
	N = 42
	fmt.Println(N)
	O := int(N)
	fmt.Println(O)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func convertStringToBytes() {

	var input string = "HELLO SRINATH"
	inputbytes := []byte(input)
	fmt.Println(inputbytes)

	for i, v := range input {
		fmt.Printf("at index position %d value is %v \n ", i, v)
	}
}
