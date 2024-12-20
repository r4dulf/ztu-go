
package main

import "fmt"

func main() {
    const Pi = 3.14159
    const (
        Red = iota
        Green
        Blue
    )
    
    fmt.Println("Value of Pi:", Pi)
    fmt.Println("Colors:", Red, Green, Blue)
}
