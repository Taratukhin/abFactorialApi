package functions

import "testing"

func TestFactorial(t *testing.T) {
	data := []struct {
		arg, res uint64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{14, 87178291200},
		{20, 2432902008176640000},
	}
	for _, d := range data {
		if got := Factorial(d.arg); got != d.res {
			t.Errorf("Factorial(%d) == %d, want %d", d.arg, got, d.res)
		}
	}
}
