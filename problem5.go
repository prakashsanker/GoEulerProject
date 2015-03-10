package main
import (
	"fmt"
)

func main() {
	for i := 20; ;i++ {
		isEvenlyDivisible := true
		for j:= 1; j < 21; j++ {
			if i % j != 0 {
				isEvenlyDivisible = false
				break
			}
		}
		if isEvenlyDivisible {
			fmt.Println(i)
			break
		}

	}

}