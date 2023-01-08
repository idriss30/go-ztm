/*
iota keyword can be used to assign integers to constants;
values can be skipped by using an underscore or blank identifier (_)
use a receiver function to more easily work with constants
*/

//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	Add = iota
	Sub
	Mul
	Div
)

type Operation int

func (operation Operation) calculate(number1, number2 int) int {
	switch operation {
	case Add:
		return number1 + number2

	case Sub:
		return number1 - number2
	case Mul:
		return number1 * number2
	case Div:
		return number1 / number2
	}
	panic("unhandle operation")
}

func main() {
	add := Operation(Add)
	sub := Operation(Sub)
	mul := Operation(Mul)
	div := Operation(Div)
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
