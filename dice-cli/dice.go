package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice(sides int) int {
	return rand.Intn(sides) + 1 // randIntn starts with 0 and a dice doesn't have a 0 number so we start with one
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	dices, sides := 3, 6
	rolls := 3

	for roll := 1; roll <= rolls; roll++ {
		sum := 0
		for d := 1; d <= dices; d++ {
			rolled := rollDice(sides)
			sum += rolled
			fmt.Println("roll #", roll, "dice #", d, ":", rolled)
		}
		fmt.Println("total rolled", sum)
		switch sum := sum; {
		case sum == 2 && dices == 2:
			fmt.Println("snake eyes")
		case sum == 7:
			fmt.Println("lucky 7")
		case sum%2 == 0:
			fmt.Println("even number")
		case sum%2 == 1:
			fmt.Println("odd number")
		}
	}

}
