/*
receiver functions provide the "dot" notation for structs
create more convenients APIS
pointer receiver can modify a struct

Value receivers cannot modify a struct
Common to use pointer receivers
Similar to modifyinf a class variable in other languages

*/

package main

import "fmt"

//tracking parking spaces

type Space struct {
	occupied bool;
}

type ParkinLot struct {
	spaces []Space
}


// regular function to compare

func occupySpace(lot *ParkinLot, spaceNumber int){
	lot.spaces[spaceNumber - 1].occupied = true;

}

// function receiver
// will allow use of dot to access the function
func (lot *ParkinLot) occupySpace2(spaceNumber int){
	lot.spaces[spaceNumber -1].occupied = true;
}

func (lot *ParkinLot) freeSpace(spaceNumber int){
	lot.spaces[spaceNumber-1].occupied = false;
}



func main(){
	lot:= ParkinLot{spaces: make([]Space, 4)}
	fmt.Println("initial parking spaces status")
	fmt.Println(lot)
	//regular function
    occupySpace(&lot, 1);
    // receiver function
	lot.occupySpace2(2)
    fmt.Println("after occupying two spaces");
	fmt.Println(lot);
	
	lot.freeSpace(1);
	fmt.Println("after freeing the first parking spot");
	fmt.Println(lot)


}
