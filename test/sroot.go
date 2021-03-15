package main

import (
	"flag"
	"fmt"
)

func main() {
	x := flag.Float64("Number", 0, "int")
	var test float64

	flag.Parse()
	test = *x
	srootCall(test)
}
func srootCall(x float64) {
	sroot := 1.0
	for roop := 1; roop <= 10; roop++ {
		presroot := sroot
		sroot = sroot - (sroot*sroot-x)/(2*sroot)
		//fmt.Println("A guess for square root is  " + strconv.FormatFloat(sroot, 'f', 15, 64))
		if presroot == sroot {
			break
		}
		fmt.Println("A guess for square root is  ", sroot)
	}
}

/*回答

package main

import "fmt"

func sqrt(num float64) float64 {
    currguess := 1.0
    prevguess := 0.0

    for count := 1; count <= 10; count++ {
        prevguess = currguess
        currguess = prevguess - (prevguess*prevguess-num)/(2*prevguess)
        if currguess == prevguess {
            break
        }
        fmt.Println("A guess for square root is ", currguess)
    }
    return currguess
}

func main() {
    var num float64 = 25
    fmt.Println("Square root is:", sqrt(num))
}
*/
