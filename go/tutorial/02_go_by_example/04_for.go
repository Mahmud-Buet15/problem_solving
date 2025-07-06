package main

import "fmt"


func main() {

	//The most basic type, with a single condition.
	i :=1
	for i<=3 {
		fmt.Println(i)
		i=i+1
	}
	fmt.Println("---------------------")

	//A classic initial/condition/after for loop.
	for j:=0; j<3; j++ {
		fmt.Println(j)
	}
	fmt.Println("---------------------")

	//Another way of accomplishing the basic “do this N times” iteration is range over an integer.
	for i:=range 3 {
		fmt.Println("range", i)
	}
	fmt.Println("---------------------")
	
	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	} 
	fmt.Println("---------------------")

	//You can also continue to the next iteration of the loop.
	for n:=range 6 {
		if n%2==0 {   //skipping if the number is even
			continue
		}
		fmt.Println(n)
	}
}



// In this code:

// 1. Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
// 2. Import the popular fmt and math packages
// 3. Implement a main function 