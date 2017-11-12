package dmath

import "math/big"

// Calculates the factorial of a number n
// and return the number calculated as
// a big.Int pointer
func Factorial(n int64) *big.Int {
    if n == 0 {
        return big.NewInt(1)
    }

    x := big.NewInt(1)
    return x.MulRange(1, n)
}
