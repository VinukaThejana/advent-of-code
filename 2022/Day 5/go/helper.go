package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Data - struct for data
type Data struct {
	Count int
	From  int
	To    int
}

func getData(line string) (data Data) {
	raw := strings.Split(line, " ")
	var err error

	data.Count, err = strconv.Atoi(raw[1])
	data.From, err = strconv.Atoi(raw[3])
	data.To, err = strconv.Atoi(raw[5])

	if err != nil {
		panic(err)
	}

	return data
}

func liftCratesPart1(stack []string, buf []string, count int) ([]string, []string) {
	tmp := make([]string, len(stack))
	copy(tmp, stack)
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	// Copy the crates until the count
	copy(buf, tmp[:count])

	// Remove the data from the stack
	for i := 0; i < count; i++ {
		stack = stack[:len(stack)-1]
	}

	return stack, buf
}

func liftCratesPart2(stack []string, buf []string, count int) ([]string, []string) {
	tmp := make([]string, len(stack))
	copy(tmp, stack)
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	// Copy the crates until the count
	copy(buf, tmp[:count])

	// Remove the data from the stack
	for i := 0; i < count; i++ {
		stack = stack[:len(stack)-1]
	}

	// Reverse the buff to implement the new crane functionality
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return stack, buf
}

func readFile() *bufio.Scanner {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}
