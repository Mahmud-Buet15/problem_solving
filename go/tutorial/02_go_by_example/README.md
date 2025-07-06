# Tutorial: Get started with Go
- Link: https://gobyexample.com/




## Code Explanations
### 06_switch.go 
```go
whatAmI := func(i interface{}) {     
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
```
#### 1. Variable Declaration with Anonymous Function
```go
whatAmI := func(i interface{}) {
    // function body
}
```
- `whatAmI` is a variable that holds an anonymous function (also called a lambda or function literal)
- The function takes one parameter `i` of type `interface{}` (the empty interface)
- In Go, `interface{}` is a special type that **can hold values of any other type**, similar to Object in Java or any in TypeScript

#### 2. Type Switch
```go
switch t := i.(type) {
    // cases
}
```
This is a **type switch**, a powerful Go feature for inspecting interface values:
- `i.(type)` is special syntax that can only be used in type switches
- It checks the concrete type (the underlying type) of the interface value i
- t receives the value of i but with its concrete type

#### 3. Case Statements
```go
case bool:
    fmt.Println("I'm a bool")
case int:
    fmt.Println("I'm an int")
default:
    fmt.Printf("Don't know type %T\n", t)
```
The switch has **three branches**:
1. If `i` is a `bool`, it prints "I'm a bool"
2. If `i` is an `int`, it prints "I'm an int"
3. For any other type (`default`), it prints the type using `%T` format verb