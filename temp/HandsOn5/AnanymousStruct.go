package main

import "fmt"

func main(){

	createAnanymousStruct()
}

func createAnanymousStruct(){

	PNR := struct {
		isGroup bool
		officeId  string
		carrierCode  string
		boardPoint string
		offPoint string

	}{true, "MA8X01RI", "8X", "NCE", "LHR"}

	fmt.Println(PNR.isGroup)
	fmt.Println(PNR.boardPoint)
	fmt.Println(PNR.officeId)
	fmt.Println(PNR.offPoint)
	fmt.Println(PNR.carrierCode)
}


