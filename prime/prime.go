package prime

import (
	"fmt"
	"math"
)

// Is test if an integer is a prime.
// It returns (true, nil) if prime, (false, error) otherwise.
// Prime number is an integer greater than 2 and divisor only 1 and itself.
func Is(num int) (bool, error) {
	min := 2
	if num < min {
		return false, fmt.Errorf("less than %d", min)
	}
	fnum := float64(num)
	sqrt := math.Sqrt(fnum)
	for i := float64(min); i <= sqrt; i++ {
		if math.Mod(fnum, i) == 0 {
			return false, fmt.Errorf("%0.0f is divisor of %d", i, num)
		}
	}
	return true,nil
}

// First get the first n primes.
// It returns slice of int
func First(num int) []int {
	x, y := 0, 0
	primes := make([]int, 0)
	for y < num {
		x++
		if b,_ := Is(x); b {
			primes = append(primes, x)
			y++
		}
	}
	return primes
}

// Between get the primes between given start and end numbers.
// It returns slice of int
func Between(start, end int) []int {
	primes := make([]int, 0)
	for i := start; i <= end; i++ {
		if b,_ := Is(i); b {
			primes = append(primes, i)
		}
	}

	return primes
}
