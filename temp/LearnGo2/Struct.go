package main

import (
	"fmt"
)

func main(){

	createStruct()
	createAnanymousStruct()
}

func createStruct(){

	type Address struct {
		streetName string
		houseNumber int
		city string
		state string
		country string

	}
	type person struct {
		firstName string
		lastName string
		Age int
		Address
}

	P1 := person{
		firstName:"SRINATH",
		lastName:"KAVURI",
		Age:25,
		Address: Address{"MunasubStreet", 123, "Gurazala", "AP", "INDIA"},
	}
	P2 := person{
		firstName:"SRINATH",
		lastName:"KAVURI",
		Age:26,
		Address: Address{"MunasubStreet", 123, "Gurazala", "AP", "INDIA"},}

	fmt.Println(P1)
	fmt.Println(P2)

	fmt.Println(P1.firstName)
	fmt.Println(P1.Age)
	fmt.Println(P1.lastName)
	fmt.Println(P1.Address)
	fmt.Println(P1.Address.houseNumber)
	fmt.Println(P1.state)

	if (P1.firstName == P2.firstName) {
		fmt.Println("Both firstNames are equal")
	}else {
		fmt.Println("Both persons firstnames are not equal")
	}


}

func createAnanymousStruct(){

	emp := struct {
		role string
		department string
		salary int
	}{"Developer", "IMPs", 2000}

	fmt.Println(emp)
}