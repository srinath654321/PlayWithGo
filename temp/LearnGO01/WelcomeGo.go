package main

import "fmt"

var y = 43

// DECLARE A varible with identifer z and that is of type "int"
// default value of int is 0
var a int
var b string
var c float32
var d float64
var e bool
var globalUserName string = "srinath kavuri"
var floatvalue = 45.05
var x int = 68

type SRINATH int

var z SRINATH = 76

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	enter()
	sayHelloToUser()
	useShortDeclaration()
	sayHelloToGo()
	loop()
	bar()
	defaults()
	printTypes()
	convertType()
	useIOTA()
}

func sayHelloToUser() {
	fmt.Println("HELLO ", globalUserName)
}

func sayHelloToGo() {
	fmt.Println("Hello welcome to go language learning")
}

func bar() {
	fmt.Println(`GO "exit the main function"`)
}

func enter() {
	fmt.Println("enter the main function")
}

func defaults() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func loop() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("number divisible by 2 ", i)
		}
	}
}

func useShortDeclaration() {
	//Short declaration operator (:=) Declare variable and assign value at the same time
	valueX := 9
	valueY := 10
	valueZ := valueX + valueY
	fmt.Println(valueZ)
}

func printTypes() {
	fmt.Println(floatvalue)
	fmt.Println("say good morning to %s ", globalUserName)
	stringOutput := fmt.Sprintf("say good evening to %s ", globalUserName)
	fmt.Println(stringOutput)
	fmt.Printf("%T\t%b\t%x\n", floatvalue, floatvalue, floatvalue)
}

func convertType() {
	type NEWTYPE int
	var m NEWTYPE = 100
	n := 56

	n = int(m)
	fmt.Println(n)

	x = int(z)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)

}

func useIOTA() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
