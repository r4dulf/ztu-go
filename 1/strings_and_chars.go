
package main

import "fmt"

func main() {
    char := 'A'
    str := "Hello, Go!"
    
    fmt.Printf("Character: %c, ASCII code: %d\n", char, char)
    fmt.Println("String length:", len(str))
    fmt.Println("Substring:", str[7:])
}
