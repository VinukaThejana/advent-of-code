package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Limit - A struct to define the lower and upper bound of the limit
type Limit struct {
	Lower int
	Upper int
}

func extractor(str string) (limit1 Limit, limit2 Limit) {
	data := strings.FieldsFunc(str, func(r rune) bool {
		return r == ',' || r == '-' || r == ' '
	})

	var err error
	limit1.Lower, err = strconv.Atoi(data[0])
	limit1.Upper, err = strconv.Atoi(data[1])
	limit2.Lower, err = strconv.Atoi(data[2])
	limit2.Upper, err = strconv.Atoi(data[3])

	if err != nil {
		panic(err)
	}

	return limit1, limit2
}

func readFile() *bufio.Scanner {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}
