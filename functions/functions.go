package main

import (
	"fmt"
);

func userName(nameInput string){
	fmt.Println("welcome to this shit show", nameInput)
}

func returnAnyMessage()string{
	return "Best way to learn go is by doing"
}
func addThreeNumbers(num1, num2, num3 int) int{
	return num1 + num2 + num3
}

func returnAnyNumber() int{
	return 7
}

func returnAnyTwoNumbers()(int,int){
	return 56, 54;

}

func addThreeNumbersCombination(){
	firstNumber, _ := returnAnyTwoNumbers()
  fmt.Println(returnAnyNumber() + addThreeNumbers(1, 4, 5) + firstNumber  )
}

func main(){
  userName("idriss");
  fmt.Println(returnAnyMessage());
  threeNumResult := addThreeNumbers(20,40,60);
  fmt.Println("the three number result is ",threeNumResult)
  
  anyNumber := returnAnyNumber();
  fmt.Println("the number that's supposed to be returned is", anyNumber)
  numberOne, numberTwo := returnAnyTwoNumbers();
  fmt.Println("the two numbers are", numberOne, numberTwo)
   
  fmt.Println("the three numbers combination sum is"  ) 
  addThreeNumbersCombination()


}