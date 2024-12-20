
package functions

func Min(a, b, c float64) float64 {
    if a < b && a < c {
        return a
    } else if b < c {
        return b
    }
    return c
}

func Avg(a, b, c float64) float64 {
    return (a + b + c) / 3
}

func SolveLinear(a, b float64) float64 {
    if a == 0 {
        panic("a cannot be zero")
    }
    return -b / a
}
