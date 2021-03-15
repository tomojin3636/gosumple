package main

import "math"


func main(){

	const(
		StatusOk = 0
		StatusConnectionReset = 1
		StatusOtherError = 2
	)

	firstName, lastName :="John" ,"Doe"
	age := 32

	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	println(integer8, integer16, integer32, integer64)

	println(math.MaxFloat32, math.MaxFloat64)

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)

	println(firstName, lastName, age)
	println(StatusOk)
	println(StatusConnectionReset)
	println(StatusOtherError)
}

