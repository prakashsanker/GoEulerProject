package main
import(
	"fmt"
)

func main() {
	//find largest palindrome made from teh product of two three digit numbers.

	//brute force -- iterate through three digit numbers, O(n^2), multiply, check for palindrome
	//speed will be O(3n^2) = O(n^2)

	//is there a faster way of doing this?
	//if I go backwards, I can fail earlier
	//so if I know that the product is lower than what I've found so far, then I don't need to keep looking?
	//can I cache results?
	/*

	999*999
	999*998
	999*997
	999*996 -> suppose this is a palindrome, I can fail and not check everything like this 999* 995, 994, 994,...
				I can also maintain a data structure (2D array, of prior calculations made, so I can save on the calculation)
				But is the calculation even that computationally difficult? I really want to save on passes through numbers. 

				If I have a stored palindrome, if my current product is lower than that, then I can discard that pass because everything after will be lower. 
	*/
			 	//highestPalindrome := 1000 * 1000
				highestPalindrome := 0
				for i := 999; i > 100 ; i-- {
					for j := 999; j > 100; j-- {
						product := i*j
						//check if palindrome
						n := product
						reversed := 0
						for product > 0 {
							digit := product % 10
							reversed = reversed*10 + digit
							product = product/10
						}

						if n == reversed {
							if n > highestPalindrome {
								highestPalindrome = n
							}
						}
					}
				}

				fmt.Println(highestPalindrome)





}