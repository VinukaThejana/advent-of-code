package main

import (
	"unicode"
)

func sumOfPriorieties() int {
	scanner := readFile()
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)

		var duplicates []string

		first := ""
		second := ""
		for i := 0; i < length; i++ {
			if i >= len(line)/2 {
				second += string(line[i])
				continue
			}

			first += string(line[i])
		}

		for j := 0; j < length/2; j++ {
			char := string(first[j])

			for k := 0; k < length/2; k++ {
				if char == string(second[k]) {
					duplicates = append(duplicates, char)
				}
			}

		}

		char := ""
		for _, value := range duplicates {
			if value != char {
				char = value
				if unicode.IsUpper([]rune(value)[0]) {
					total += capitalLetters[value]
					continue
				}

				total += simpleLetters[value]
			}
		}

		// Reset
		first = ""
		second = ""
		duplicates = nil
	}

	return total
}
