package main

import (
	"fmt"
)

func main() {

	var primes []int

	for i := 2; i < 100; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	fmt.Println("Hello, I'm project 0")
	fmt.Println("Primes[]int ", primes)

}
