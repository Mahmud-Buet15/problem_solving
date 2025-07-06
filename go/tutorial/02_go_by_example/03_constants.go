package main

import (
    "fmt"
    "math"
)

const s string = "constant"    //const declares a constant value. A const statement can appear anywhere a var statement can.

func main() {
	fmt.Println(s)

	const n= 500000000
	fmt.Println(n)
	fmt.Println(3e20)


	const d = 3e20 / n    
	fmt.Println(d)			// Constant expressions perform arithmetic with arbitrary precision.
	fmt.Println(int64(d))   //A numeric constant has no type until it’s given one, such as by an explicit conversion.

	fmt.Println(math.Sin(n))
}



// In this code:

// 1. Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
// 2. Import the popular fmt and math packages
// 3. Implement a main function 