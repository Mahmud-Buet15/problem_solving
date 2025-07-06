// Go supports anonymous functions, which can form closures. 
// Anonymous functions are useful when you want to define a function inline without having to name it

package main

import "fmt"

// This function intSeq returns another function, which we define anonymously in the body of intSeq. 
// The returned function closes over (captures) the variable i to form a closure. This means i persists between calls to the returned function.
func intSeq() func() int {
    i := 0

    return func() int {
        i++
        return i
    }
}


func main() {

	// We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
    nextInt := intSeq()

    fmt.Println(nextInt())  // 1
    fmt.Println(nextInt())  // 2
    fmt.Println(nextInt())  // 3
	fmt.Println("----------------------")

    newInts := intSeq()
    fmt.Println(newInts())  // 1
}

// Key Concepts Demonstrated
// 1. Closures: Functions that remember values from their enclosing scope
// 2. First-class functions: Functions can return other functions
// 3. State preservation: The variable i persists between function calls
// 4. Isolation: Each closure maintains its own separate state