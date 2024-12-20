
package main

import "fmt"

func main() {
    var array = [5]int{1, 2, 3, 4, 5}
    slice := array[1:4]
    
    fmt.Println("Array elements:", array)
    fmt.Println("Slice elements:", slice)
}
