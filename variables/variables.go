package main

import (
	"fmt"
)

func main() {
	var myName string = "dalton"
	var zeroVariable float32
	const CurrentYear = "2022" // can not be reassigned like in javascript
	fmt.Println("my name is", myName)
	languageOne, languageTwo := "Go", "javascript"
	fmt.Println("I solve problems with", languageOne, "and", languageTwo)
	fmt.Println("the current year is", CurrentYear)
	fmt.Println("A zero variable is", zeroVariable)

	//with comma ok variable the second variable can be reassigned
	firstVariable, otherVariable := "john", "doe"
	fmt.Println("firstVariable is", firstVariable, "otherVariable is", otherVariable)
	secondVariable, otherVariable := "joe", "dalton"
	fmt.Println("secondVariable is", secondVariable, "secondVariable  reassignment is", otherVariable)

}
