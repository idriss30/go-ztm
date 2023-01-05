package main

import "fmt"

func printAssemblyLine(Content []string) {
	for i := 0; i < len(Content); i++ {
		item := i
		fmt.Println("the current content is", Content[item])
	}

}

func printStringWordsWithRange() {
	mySlice := []string{"hello", "Go", "language"}
	for i, element := range mySlice {
		index := i
		fmt.Printf("the word is %v and the index is %v\n", element, index)

		for _, character := range element {
			fmt.Printf("the character is %q\n", character) // if not adding the %q it will print the bite number of the rune and not the actual letter
		}
	}
}
func main() {
	var assemblyContent = []string{"carts", "convoyeurs belts", "workers"}
	printAssemblyLine(assemblyContent)
	assemblyContent = append(assemblyContent, "pulleys", "tapes")
	fmt.Println("##############################3")
	printAssemblyLine(assemblyContent)
	newAssemblyContent := assemblyContent[3:]
	fmt.Println("############")
	printAssemblyLine(newAssemblyContent)

	fmt.Println("//////////////////// printing with range")
	printStringWordsWithRange()

}
