package main

import (
	"challenges/sort"
	"fmt"
	"math/rand"
)

func generates(count int, max int) []int {
	a := make([]int, count)
	for i := count - 1; i >= 0; i-- {
		a[i] = rand.Intn(max)
	}
	return a
}

func main() {
	collection := generates(10, 100)
	fmt.Println(collection)

	asc := sort.IntAsc(collection)
	fmt.Println(asc)

	desc := sort.IntDesc(collection)
	fmt.Println(desc)
}
