//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendedCheck struct {
	checkout time.Time
	checkIn  time.Time
}
type Members struct {
	name  Name
	books map[Title]LendedCheck
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Members
	books   map[Title]BookEntry
}

func printMemberAudit(member *Members) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkout.String(), "through", returnTime)
	}

}

func printMemberFunctionValue(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}

}

func printLibraryBooks(library *Library) {
	fmt.Println("#############")
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println("###############")
}

func checkoutBook(library *Library, title Title, member *Members) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println(("book not part of library"))
		return false
	}
	if book.lended == book.total {
		fmt.Println("there is no books available to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book
	member.books[title] = LendedCheck{checkout: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Members) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println(("book not part of library"))
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true

}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Members),
	}

	library.books["Webapp in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["the little Go book"] = BookEntry{
		total:  3,
		lended: 0,
	}

	library.books["Idiomatic Go"] = BookEntry{
		total:  6,
		lended: 0,
	}
	library.books["Go bootcamp"] = BookEntry{
		total:  2,
		lended: 0,
	}

	library.members["joe"] = Members{"Joe", make(map[Title]LendedCheck)}
	library.members["dalton"] = Members{"dalton", make(map[Title]LendedCheck)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberFunctionValue(&library)

	member := library.members["dalton"]
	checkedOut := checkoutBook(&library, "Idiomatic Go", &member)
	fmt.Println("\n Check out a book")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberFunctionValue(&library)
	}
	returned := returnBook(&library, "Idiomatic Go", &member)
	fmt.Println("\n Check in a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberFunctionValue(&library)
	}
}
