package main

import (
	"fmt"
)

func main() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		switch {
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		case val < 0:
			panic("Panic in g() (major)")
		default:
			fmt.Println("You entered:", val)

		}

	}

}
