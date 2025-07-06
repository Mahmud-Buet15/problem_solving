package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)   //AND logic
	fmt.Println(true || false)   //OR logic
	fmt.Println(!true)			 //NOT logic
}



// In this code:

// 1. Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
// 2. Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
// 3. Implement a main function 