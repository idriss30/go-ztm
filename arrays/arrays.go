package main

import (
	"fmt"
)

type Rooms struct {
	name    string
	cleaned bool
}

func checkCleaningLess(rooms [4]Rooms) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "has been cleaned")
		} else {
			fmt.Println(room.name, "is still dirty")
		}
	}

}

func main() {
	rooms := [4]Rooms{
		{name: "office", cleaned: true},
		{name: "reception", cleaned: false},
		{name: "operations", cleaned: false},
		{name: "warehouse", cleaned: true},
	}

	checkCleaningLess(rooms)

	rooms[0].cleaned = false
	rooms[2].cleaned = true

	checkCleaningLess(rooms)

}
