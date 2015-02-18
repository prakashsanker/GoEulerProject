// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

package main
import (
	"fmt"
)

func main() {
	//check if a number is a multiple of 3
	//check if a numebr is a multiple of 5
	sum := 0
	for i:=0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}