
package math

import "testing"

func TestMin(t *testing.T) {
    if Min(3, 1, 2) != 1 {
        t.Error("Expected 1, got", Min(3, 1, 2))
    }
}

func TestAvg(t *testing.T) {
    if Avg(3, 1, 2) != 2 {
        t.Error("Expected 2, got", Avg(3, 1, 2))
    }
}

func TestSolveLinear(t *testing.T) {
    if SolveLinear(3, 2) != -2.0/3 {
        t.Error("Expected -2/3, got", SolveLinear(3, 2))
    }
}
