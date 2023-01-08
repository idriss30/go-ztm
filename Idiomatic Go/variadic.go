/* when not knowing how many arguments a function can take
use variadic function */

package main

// calculate sum of numbers without knowing the number of arguments
import "fmt"

func sum(numbers ...uint) uint {
	var sum uint
	for _, number := range numbers {
		sum += number
	}
	return sum

}

func main() {
	a := []uint{1, 2, 4}
	b := []uint{3, 5, 6}

	all := append(a, b...)
	fmt.Println(all)
	sumValue := sum(1,4,5);
	secondSumValue := sum(all...)
	fmt.Println("the sum is", sumValue);
	fmt.Println("second sum value", secondSumValue)

}
