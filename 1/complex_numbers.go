
package main

import "fmt"

func main() {
    var complexNum complex64 = 1 + 2i
    fmt.Println("Complex number:", complexNum)
    fmt.Println("Real part:", real(complexNum))
    fmt.Println("Imaginary part:", imag(complexNum))
}
