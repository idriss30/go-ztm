package main

import "fmt"

func main() {
	mySlice := []int{}
    slice2 := []int{7,8,9}
	mySlice = append(mySlice, 1, 2, 3)
	fmt.Printf("the value of my slice is %v\n", mySlice,);
	mySlice = append(mySlice, slice2...) // combine second slice with triple dot
	fmt.Println("the new value of my slice is", mySlice)


	// slice into array 
	mystringArray := [4]string{"john", "joe", "williams", "tobi"};
	stringSlice := mystringArray[1:3];
	secondStringSlice := mystringArray[:]
	fmt.Println(stringSlice)
	secondStringSlice = append(secondStringSlice, "luffy", "roronoa", "chopper");

	fmt.Println("here is the second slice", secondStringSlice);
	fmt.Println("the initial array is still present", mystringArray)

}
