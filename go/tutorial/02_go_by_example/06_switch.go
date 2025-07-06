package main

import (
    "fmt"
    "time"
)

func main() {

    // Here’s a basic switch.
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:                 // checks if i==1
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
    fmt.Println("-------------------")

    // You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:     // checks if time.Now().Weekday()==time.Saturday or time.Sunday
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
    fmt.Println("-------------------")

    // switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
    t := time.Now()
    switch {
    case t.Hour() < 12:                 // checks if time.Now().Hour()<12
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
    fmt.Println("-------------------")


    // A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.
    whatAmI := func(i interface{}) {      //details in README.md
        switch t := i.(type) {            //This is a type switch
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}




// In this code:

// 1. Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
// 2. Import the popular fmt and math packages
// 3. Implement a main function 

