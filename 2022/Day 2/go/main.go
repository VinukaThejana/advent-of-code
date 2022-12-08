package main

import "fmt"

var total = 0

func main() {
	getTotal()
	fmt.Println(total)

	// Reset the total for the second challenge
	total = 0

	getNewTotal()
	fmt.Println(total)
}
