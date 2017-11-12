package dmath


import "math"

// Return whether a number is
// a prime number
func IsPrime(n int) bool {
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}
