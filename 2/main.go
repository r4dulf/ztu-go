
package main

import (
    "fmt"
    "myapp/functions"
)

func main() {
    fmt.Println("Min:", functions.Min(3, 1, 2))
    fmt.Println("Avg:", functions.Avg(3, 1, 2))
    fmt.Println("Solution of 3x + 2 = 0:", functions.SolveLinear(3, 2))
}
