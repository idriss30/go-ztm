//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Items struct {
	name     string
	active   bool
	inactive bool
}

func activateTag(item *Items) {
	item.inactive = false
	item.active = true
}

func deactivateTag(item *Items) {
	item.inactive = true
	item.active = false
}

func checkout(items []Items) {
	for i :=range items{
        deactivateTag(&items[i])
	}
}

func main() {
	items := []Items{
		{name: "EmployeeTag", active: true, inactive: false},
		{name: "VisitorTag", active: true, inactive: false},
		{name: "restaurantTag", active: true, inactive: false},
		{name: "OwnerTag", active: true, inactive: false},
	}
	fmt.Println("items before changes")
	fmt.Println(items)

	deactivateTag(&items[2])
	fmt.Println("items after item 3 got deactivated changes")
	fmt.Println(items)
	activateTag(&items[2])
	fmt.Println("items should be back to initial")
	fmt.Println(items)
	checkout(items);
	fmt.Println("final everything deactivated");
	fmt.Println(items)

}

