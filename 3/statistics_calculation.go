
package main

import (
    "fmt"
    "math"
)

func calculateStats(numbers []int) (float64, float64, float64) {
    n := float64(len(numbers))
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    mean := float64(sum) / n

    varianceSum := 0.0
    for _, num := range numbers {
        varianceSum += math.Pow(float64(num)-mean, 2)
    }
    variance := varianceSum / n
    stdDev := math.Sqrt(variance)

    return mean, variance, stdDev
}

func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    mean, variance, stdDev := calculateStats(data)
    fmt.Printf("Mean: %.2f, Variance: %.2f, StdDev: %.2f\n", mean, variance, stdDev)
}
