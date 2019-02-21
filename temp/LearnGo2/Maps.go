package main

import "fmt"

func main(){

	createMap()
}

func createMap(){

	StudentAgeMap := map[string]int{
		"SRINATH" : 25,
		"SARANYA" : 23,
		"ANASUYA" : 49,
		"BHARATH" : 52,
	}
	fmt.Println(StudentAgeMap)

	fmt.Println(StudentAgeMap["SRINATH"])

	if v, ok := StudentAgeMap["LAVU"] ; ok {
		fmt.Println("PRINT THE VALUE ", v)
	}else if v, ok :=StudentAgeMap["ANASUYA"]; ok{
		fmt.Println("PRINT TO COSOLE ", v)
	}
	StudentAgeMap["ANUSHA"] = 26
	StudentAgeMap["DIVYA"] = 27

	for key,value := range StudentAgeMap {
		fmt.Printf("Key : %s, value : %d \n", key, value)
	}


	if value, ok := StudentAgeMap["SRINATH"]; ok{
		fmt.Println(value)
		delete(StudentAgeMap, "SRINATH")
	}
	fmt.Println(StudentAgeMap)
}
