package functions

func Factorial(n uint64) (result uint64) { // function is not recursive, I don't want to use stack
	result = 1 // if n==0 { Factorial(0)==1 }
	for i := uint64(1); i <= n; i++ {
		result *= i
	}
	return result
}
