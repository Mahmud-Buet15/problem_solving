package main

import (
    "fmt"
    "maps"
)

func main() {

    // To create an empty map, use the builtin make: make(map[key-type]val-type).This is like python dictionary
    m := make(map[string]int)
    fmt.Println("map:", m)
    fmt.Println("----------------------")

    // Set key/value pairs using typical name[key] = val syntax.
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m)
    fmt.Println("----------------------")

    // Get a value for a key with name[key].
    v1 := m["k1"]
    fmt.Println("v1:", v1)
    fmt.Println("----------------------")

    // If the key doesn’t exist, the zero value of the value type is returned.
    v3 := m["k3"]
    fmt.Println("v3:", v3)
    fmt.Println("----------------------")

    // The builtin delete removes key/value pairs from a map.
    delete(m, "k2")
    fmt.Println("map:", m)
    fmt.Println("----------------------")

    // To remove all key/value pairs from a map, use the clear builtin.
    clear(m)
    fmt.Println("map:", m)
    fmt.Println("----------------------")

    // The optional second return value when getting a value from a map indicates if the key was present in the map.
    // This can be used to disambiguate between missing keys and keys with zero values like 0 or "". 
    // Here we didn’t need the value itself, so we ignored it with the blank identifier _.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)                //false indicates missing key
    fmt.Println("----------------------")

    // You can also declare and initialize a new map in the same line with this syntax.
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
    fmt.Println("----------------------")

    // The maps package contains a number of useful utility functions for maps.
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}