package main

import "fmt"

func main() {

    //Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. 
    //By default an array is zero-valued, which for ints means 0s.
    var a [5]int
    fmt.Println("emp:", a)
    fmt.Println("len:", len(a))
    fmt.Println("----------------------")

    //We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    fmt.Println("----------------------")

    //Use this syntax to declare and initialize an array in one line.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
    fmt.Println("----------------------")
    
    //You can also have the compiler count the number of elements for you with ...
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
    fmt.Println("----------------------")

    //If you specify the index with :, the elements in between will be zeroed.
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
    fmt.Println("----------------------")

    //Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
    var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
    fmt.Println("----------------------")

    //You can create and initialize multi-dimensional arrays at once too.
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}



// In this code:

// 1. Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
// 2. Import the popular fmt and math packages
// 3. Implement a main function 

