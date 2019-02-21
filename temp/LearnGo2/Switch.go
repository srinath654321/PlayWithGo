package main

import (
	"fmt"
	"time"
)

func main(){

	time.Now().Weekday()
	switch "WEEKDAY" {

	case "MONDAY", "WEEKDAY":
		fmt.Println("Today is monday a week day")
		fallthrough
	case "TUESDAY":
		fmt.Println("TOday is tuesday a week day")
	default:
		fmt.Println("print default value")
		}
}
