// Advent of code Day 4
package main

import "fmt"

func main() {
	totalIntersectionCount := totalIntersection()
	fmt.Println(totalIntersectionCount)

	partialIntersectionsCount := partialIntersections()
	fmt.Println(partialIntersectionsCount)
}
