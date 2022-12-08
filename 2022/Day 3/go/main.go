// Advent of code
package main

import "fmt"

func main() {
	// Challenge 1
	// Get the total
	totalPriorities := sumOfPriorieties()
	fmt.Println(totalPriorities)

	// Challenge 2
	// Get the total with groups
	totalWithGroups := groupedRustacks()
	fmt.Println(totalWithGroups)
}
