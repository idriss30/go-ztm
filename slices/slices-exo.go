package main

import "fmt"


func printAssemblyLine(Content []string) {
	for i := 0; i < len(Content); i++ {
		item := i
		fmt.Println("the current content is", Content[item])
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

}
