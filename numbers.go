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

// Returns the greatest common devisor
// of two ints x and y
func Gcd(x, y int) int {
    r := x % y
    c := x

    for r != 0 {
        c = r
        r = y % r
        y = c
    }

    return c
}
