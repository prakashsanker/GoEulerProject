// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

package main
import (
	"fmt"
)

func main() {
	//how to find prime factors?
	//
	fmt.Println(getLargestPrimeFactor(600851475143));
}



func getLargestPrimeFactor(number int64) []int64{
	//what is a prime factor?
	//a number that multiplies with another number to get the target number and is prime
	//when is a number prime - when it is only divisible by itself and one. 

	//find all factors
	//decide which ones are prime.

	factorStore := make([]int64,0)
	knownPrimes := make()
	for i := int64(2); i < number + int64(1); i = i+2 {
		if number%i == int64(0) {
			//figure out what it multiplies by
			isPrime := true
			for j := int64(2); j < i; j++ {
				if i%j == int64(0) {
					if j != int64(1) {
						//it is not prime
						isPrime = false
					}
				}
			}

			if isPrime {
				factorStore = append(factorStore, i)
			}
		}
	}

	return factorStore



	//that should give us a data structure with the factors for the number
	//now go through the numbers in that list and figure out which ones are prime.
	//how to check if a number is prime?
	//go through 

}