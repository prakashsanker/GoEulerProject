package main
import (
	"fmt"
)

//It's clear that via 'math' that (a + b + c + d...)^2 = a^2 + b ^2 + c^2 + d^2 + 2ac + 2ad + 2db + 2cd + 2bc + 2ab ....
//So the sum of squares is inside that calculation
//really, we want to calculate 2ab + 2ac + 2ad...

func main() {
	sum := 0
	for i := 1; i < 101; i++ {
		for j:= 1; j < 101; j++ {
			if i != j {
				sum += i*j
			}
		}
	}
	fmt.Println(sum)
}