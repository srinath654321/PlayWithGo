package main

import "fmt"

func main(){
    // do not use arrays in go use slices
	declareArray()
	createSlice()
	createMake()
	createMultiDimensionalSlice()
}

func declareArray(){

	var x[5] int
	fmt.Println(x)
	fmt.Println(len(x))
}

func createSlice(){

	// composite literal type{values}
	M := []int{5,6,7,8}
	fmt.Println(len(M))
	fmt.Println(M)
	fmt.Println(M[1])
	fmt.Println("")

	for i, v:= range M {
		fmt.Println(i, v)
	}

	fmt.Println(M[:])
	fmt.Println(M[2:3])

	fmt.Println("")

	for i:=0; i<len(M); i++ {

		fmt.Println(i, M[i])
    }

	M = append(M, 9, 10, 11, 12)
	fmt.Println(M)
    fmt.Println("")

	N := []int{34,35,36,37}
	M = append(M,N...)
	fmt.Println(M)

	fmt.Println("")

	M = append(M[:0], M[1:]...)

	fmt.Println(M)


}

func createMake(){

	 initialArray := make([]int, 20, 300)
	 fmt.Println(initialArray)
	 fmt.Println(len(initialArray))
	 fmt.Println(cap(initialArray))
	 initialArray = append(initialArray, 450)
	 fmt.Println(initialArray)
}

func createMultiDimensionalSlice(){

	aj := []string {"srinath", "saranya", "anasuya", "bharath"}
	bj := []string {"samba", "kiran" ,"bhavitha", "vanshika"}
	fmt.Println(aj)
	fmt.Println(bj)
	cj := [][]string{aj, bj}
	fmt.Println(cj)

	for index, value := range cj{
		fmt.Println(index, value)
	}

	for i,v := range aj{
		fmt.Println(i,v)
	}
}