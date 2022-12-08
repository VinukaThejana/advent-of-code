package main

import (
	"bufio"
	"os"
)

type MatchState string

const (
	W MatchState = "W"
	D            = "D"
	L            = "L"
)

type ObjectState string

const (
	R ObjectState = "R"
	P             = "P"
	S             = "S"
)

type OpponentResponses string

const (
	// Rock
	A OpponentResponses = "A"
	// Paper
	B = "B"
	// Scissor
	C = "C"
)

type YourResponses string

const (
	// Rock
	X YourResponses = "X"
	// Paper
	Y = "Y"
	// Scissor
	Z = "Z"
)

// Define a hashmap to store the match scores
var match = map[string]int{
	string(W): 6,
	string(D): 3,
	string(L): 0,
}

// Define a hashmap to store the object scores
var object = map[string]int{
	string(R): 1,
	string(P): 2,
	string(S): 3,
}

// Match scores
func getMatchScores(state MatchState) int {
	return match[string(state)]
}

// Object Scores
func getObjectScores(state ObjectState) int {
	return object[string(state)]
}

// Read the file
func readFile() *bufio.Scanner {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}
