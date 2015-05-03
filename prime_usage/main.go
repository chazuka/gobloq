package main

import (
	"challenges/prime"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	// get the first 25 prime numbers
	num := 25
	fmt.Println(fmt.Sprintf("The first %d primes are:", num))
	firstNumPrimes := prime.First(num)
	fmt.Println(firstNumPrimes)

	// get prime numbers between 1 and 100
	start, end := 1, 100
	primesBetween := prime.Between(start, end)
	fmt.Println(fmt.Sprintf("There are %d primes between %d and %d", len(primesBetween), start, end))
	fmt.Println(primesBetween)

	// random is prime check
	for i := 0; i < 10; i++ {
		n := rand.Intn(1000)
		b, err := prime.Is(n)
		m := (map[bool]string{
			true:  "is a prime",
			false: fmt.Sprintf("is not a prime! %s", err),
		})[b]
		fmt.Println(strconv.Itoa(n), m)
	}

}
