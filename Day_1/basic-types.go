package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	/**
	Basic Types are:
	- bool
	- string
	- int, int8, int16, int32, int64
	- uint, uint8, uint16, uint32, uint64, uintptr //unsigned integer
	- byte // alias of int8
	- rune // alias of int32, it represents unicode code point
	- float32, float64
	- complex64, complex 128
	*/

	/**
	int, uint, uintptr usually comes in size of 32bits of 32-bit system and 64bits on 64-bit system
	*/
	var (
		theBool    bool       = true
		theString  string     = "now you see me"
		theInt16   uint16     = 1<<16 - 1
		theInt8    byte       = 127
		theComplex complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", theBool, theBool)
	fmt.Printf("Type: %T Value: %v\n", theString, theString)
	fmt.Printf("Type: %T Value: %v\n", theInt16, theInt16)
	fmt.Printf("Type: %T Value: %v\n", theInt8, theInt8)
	fmt.Printf("Type: %T Value: %v\n", theComplex, theComplex)

	/**
	Type can be converted using expression like this T(v)
	it means, it converts the value of v into the type of T
	*/

	var toFloat64 = float64(theInt16)
	fmt.Printf("Type: %T Value: %v\n", toFloat64, toFloat64)

}
