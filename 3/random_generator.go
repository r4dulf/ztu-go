
package main

import "fmt"

func generateRandom(seed, a, c, m, n int) []int {
    results := make([]int, n)
    x := seed
    for i := 0; i < n; i++ {
        x = (a*x + c) % m
        results[i] = x
    }
    return results
}

func main() {
    randomNumbers := generateRandom(1, 1103515245, 12345, 1<<31, 10)
    fmt.Println("Random numbers:", randomNumbers)
}
