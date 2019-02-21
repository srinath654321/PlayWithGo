package main

import "fmt"

func main(){

	createPerson()
}

func createPerson(){

	type person struct {

		firstName string
		lastName string
		favouriteIceCreamFlavours []string
	}

	Person1 := person{"SRIANTH", "KAVURI", []string{"Vennila", "Straberry", "chocolate"}}
	Person2 := person{"RAM", "NANDAMURI", []string{"BANANA", "MansionHouse", "Carrot"}}

	fmt.Println(Person1)
	fmt.Println(Person1.firstName)
	fmt.Println(Person1.lastName)
	flavours1 := Person1.favouriteIceCreamFlavours
	for _,flavors := range flavours1 {
		fmt.Println(flavors)
	}
	fmt.Println(Person2)

	for _, flavours := range Person2.favouriteIceCreamFlavours {
		fmt.Println(flavours)
	}

personsMap :=	map[string]person {
		                "KAVURI" : Person1,
		                 "NALLA" : Person2,
	               }
for _, value := range personsMap {
	fmt.Println(value.favouriteIceCreamFlavours)
	for _, flavours := range value.favouriteIceCreamFlavours {
		fmt.Println(flavours)
	}
	fmt.Println(value.lastName)
	fmt.Println(value.firstName)
}
}