package main

import "fmt"

func main()  {

	createVehicle()
}


func createVehicle(){

	type Vehicle struct {
		doors int
		color string
	}

	type Truck struct {
		Vehicle
		fourWheels bool
	}

	type Sedan struct {
		Vehicle
		luxury  bool
	}

	TOYOTA_TACOMA := Truck{Vehicle{4, "persianPerl"}, true}
	TOYOTA_CAMRY  := Sedan{Vehicle{4, "GRAY"}, true}

	fmt.Println(TOYOTA_TACOMA.color)
	fmt.Println(TOYOTA_TACOMA.doors)
	fmt.Println(TOYOTA_TACOMA.fourWheels)
	fmt.Println(TOYOTA_CAMRY.color)
	fmt.Println(TOYOTA_CAMRY.doors)
	fmt.Println(TOYOTA_CAMRY.luxury)
}