package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Alway returns a fixed number

	/*
		Intn returns, as an int,
		a non-negative pseudo-random number in the half-open interval [0,n)
		from the default Source. It panics if n <= 0.
	*/

	fmt.Println(rand.Intn(10))

	// returns new number each random

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println(r.Intn(10))

}
