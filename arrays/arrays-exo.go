package main

import "fmt"

// create a program that store a shopping list and print out information about the list.

type Product struct {
	name  string
	price int
}

func getTotal(list [4]Product) (int, int) {
	total := 0
	totalItem := 0
	for i := 0; i < len(list); i++ {
		item := list[i]
		total += item.price

		if item.name != ""{
			totalItem ++
		}
	}
	return total, totalItem
}

func printListInfo(list [4]Product) {
	totalCost, totalItems := getTotal(list);
	fmt.Println("the total number of items", totalItems);
	fmt.Println("the total cost of the shopping list is", totalCost)
	fmt.Println("the last item on the list", list[len(list)-1])
	
}

func main() {
	Product := [4]Product{
		{name: "tomato", price: 10},
		{name: "onions", price: 9},
		{name: "potatoes", price: 15},
	}
	printListInfo(Product);

	Product[3].name = "rice";
	Product[3].price = 30;
	printListInfo(Product)
}
