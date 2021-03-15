package main

import (
	"fmt"
)

func main() {
	for roop := 1; roop <= 100; roop++ {
		switch {
		case roop%3 == 0 && roop%5 == 0:
			fmt.Println("FizzBuzz")
		case roop%3 == 0:
			fmt.Println("Fizz")
		case roop%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(roop)
		}
	}

}

/*回答
package main

import (
    "fmt"
    "strconv"
)

func fizzbuzz(num int) string {
    switch {
    case num%15 == 0:
        return "FizzBuzz"
    case num%3 == 0:
        return "Fizz"
    case num%5 == 0:
        return "Buzz"
    }
    return strconv.Itoa(num)
}

func main() {
    for num := 1; num <= 100; num++ {
        fmt.Println(fizzbuzz(num))
    }
}

*/
