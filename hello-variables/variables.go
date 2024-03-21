package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    // variable types are inferred from asignment
    var d = true
    fmt.Println(d)

    // declared variables but no initialized are given a empty value accordion to type aka zero-valued
    var e int
    fmt.Println(e)

    // for booleans zero-value is false
    var ee bool
    fmt.Println(ee)

    f := "apple" // short hand for declaration and initialization only available inside function scope
    fmt.Println(f)
}
