package main

import "fmt"

func main() {
	cals := getTheElfWithMostcalories()
	fmt.Println(cals)

	totalTopThreeCals := getTheTotalCalsOfTheTopThreeElfs()
	fmt.Println(totalTopThreeCals)
}
