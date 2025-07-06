package main

import "fmt"

func main() {
	var a="initial"     // declaring a variable
	fmt.Println(a)

	var b,c int = 1,2     // declaring multiple variables
	fmt.Println(b,c)

	var d = true         // Go will infer the type of initialized variables.
	fmt.Println(d)

	var e int  
	fmt.Println(e)     // if nothing is assigned, 0 is assigned by default

	var f string  
	fmt.Println(f)     // if nothing is assigned, blank string is assigned by default

	g := "apple"      // := syntax is shorthand for declaring and initializing a variable
	fmt.Println(g)
}



// In this code:

// 1. Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
// 2. Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
// 3. Implement a main function