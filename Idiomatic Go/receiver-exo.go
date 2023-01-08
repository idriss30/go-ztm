//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health,
//  - Energy,
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	Name   string
	Health uint
	Energy uint
}

//creating receiver function

func (player *Player) modifyHealth(number uint) {
	fmt.Println("the current health status of the player is ", player.Health)
	player.Health = number
	fmt.Println("health after modification is", player.Health)
}

func (player *Player) modifyEnergy(number uint) {
	fmt.Println("current energy", player.Energy)
	player.Energy = number
	fmt.Println("after modification", player.Energy)
}

func main() {
	player := Player{Name: "thomas", Energy: 200, Health: 300}
	player.modifyEnergy(100)
	player.modifyHealth(500)
}
