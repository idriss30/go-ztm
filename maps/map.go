/*
	Maps are a commonly used data structure that stores data in key value pairs

the performance is very high when the key is known
Unordered- data is stored in random order in contrast to an array
*/
package main

import "fmt"

func main() {
	myMap := make(map[string]int) // meaning string key int value it is uninitialized;
	myMap["favorite number"] = 7  // insert into Map
	fav := myMap["favorite number"]
	fmt.Println("the favorite number is ", fav)

	//deleting one item
	delete(myMap, "favorite number")

	//find item
	price, found := myMap["price"]
	if !found {
		fmt.Println("price was not found")
	} else {
		fmt.Println("price was found the value is", price)
	}
	// to iterate with the map we use the range keyword
	// you get a key and value pair instead of a index and value

}
