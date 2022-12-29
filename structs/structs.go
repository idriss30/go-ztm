package main

import (
	"fmt"
)

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	john := Passenger{"john", 2, true}
	fmt.Println("these are john details", john)
	var dalton = Passenger{Name: "dalton", TicketNumber: 1}
	fmt.Println("these are dalton details", dalton)

	var (
		april = Passenger{Name: "april", TicketNumber: 3, Boarded: true}
		avrel = Passenger{Name: "avrel", TicketNumber: 4}
	)

	fmt.Println("details for april and avrel", april, avrel)

	if april.Boarded {
		fmt.Println(april.Name, "has boarded the bus")
	}
	avrel.Boarded = true
	if avrel.Boarded {
		fmt.Println(avrel.Name, "has boarded the bus")
	}

	dalton.Boarded = true
	if dalton.Boarded {
		fmt.Println(dalton.Name, "has boarded the bus")
	}

	var lastPassenger = Passenger{}
	lastPassenger.Name = "Donatelli versace"
	lastPassenger.TicketNumber = 8
	lastPassenger.Boarded = true

	var busOne = Bus{lastPassenger}
	fmt.Println("the passenger seating in the front is", busOne.FrontSeat.Name, "\n", "with a ticket number ", busOne.FrontSeat.TicketNumber)
}
