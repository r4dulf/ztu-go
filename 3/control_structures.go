
package main

import "fmt"

func main() {
    // if-else конструкція
    a, b := 5, 10
    if a > b {
        fmt.Println("a більше за b")
    } else {
        fmt.Println("a менше або дорівнює b")
    }

    // for цикл
    for i := 1; i <= 5; i++ {
        fmt.Println("Ітерація:", i)
    }

    // switch
    number := 3
    switch number {
    case 1:
        fmt.Println("Один")
    case 2:
        fmt.Println("Два")
    default:
        fmt.Println("Інше значення")
    }
}
