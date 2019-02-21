package main

import "fmt"

func main()  {

	for i:=0; i<=10; i++ {
		if i%2 != 0 {

			if i==9 {

				continue
			}
			fmt.Println("print number less than or equal to 4 ", i)
		}
	}

	a := 4
	b := 20

	for a < b {
		a = a + 4
		fmt.Println(a)
	}

	fmt.Println("")

	for {

		if b < 10 {
			break
		}
		b--
		fmt.Println(b)

	}

	for i:=33; i<124; i++{
		fmt.Printf("%#U\n", i)
	}
}
