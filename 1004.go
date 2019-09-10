package main
 
import (
    "fmt"
)
 
func main() {

    var a, b, c int
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    c = a * b
    fmt.Println("PROD =",c)

}
