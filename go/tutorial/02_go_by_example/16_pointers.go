
// Go supports pointers, allowing you to pass references to values and records within your program.
package main

import "fmt"

// We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr. 
func zeroval(ival int) {
    ival = 0
}
// Takes an integer parameter ival (passed by value)
// Modifies the local copy of ival inside the function
// Changes don't affect the original variable outside the function


// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer. 
func zeroptr(iptr *int) {
    *iptr = 0
}
// Takes a pointer to an integer (*int)
// Dereferences the pointer with * to modify the original value
// Changes affect the original variable outside the function

func main() {
    i := 1
    fmt.Println("initial:", i)		// zeroval: 1 (unchanged)		//Passes a copy of i (value doesn't change outside)
	fmt.Println("----------------------")

    zeroval(i)
    fmt.Println("zeroval:", i)		// zeroptr: 0 (changed)			//Passes the memory address of i (value changes)
	fmt.Println("----------------------")

    zeroptr(&i)						//The &i syntax gives the memory address of i, i.e. a pointer to i.
    fmt.Println("zeroptr:", i)
	fmt.Println("----------------------")

    fmt.Println("pointer:", &i)		// prints memory address where i is stored
}



// zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.

// Key Concepts Demonstrated
// 1. Pass by Value (default in Go):
// - Functions receive copies of arguments
// - Modifications don't affect originals
// - Safe but can be inefficient for large structs

// 2. Pass by Pointer:
// - Use & to get a variable's address
// - Use * to dereference (access the pointed-to value)
// - Allows modifying original variables
// - Essential for sharing state across function boundaries

// 3. Pointer Syntax:
// - *int - type "pointer to int"
// - &i - address-of operator (gets pointer to i)
// - *iptr - dereference operator (accesses the value)