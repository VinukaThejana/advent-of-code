package main

import (
	"bufio"
	"os"
)

func readFile() *bufio.Scanner {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}
