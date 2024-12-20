
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Введіть ваше ім'я: ")
    fmt.Scan(&name)
    fmt.Print("Введіть ваш вік: ")
    fmt.Scan(&age)
    fmt.Printf("Привіт, %s! Вам %d років.\n", name, age)
}
