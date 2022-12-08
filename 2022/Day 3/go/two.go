package main

import (
	"unicode"
)

func artihmatic(i int) int {
	return 3*i - 2
}

func groupedRustacks() int {
	scanner := readFile()
	total := 0

	n := 1
	i := 1

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		line1Index := artihmatic(i)
		line2Index := artihmatic(i) + 1
		line3Index := artihmatic(i) + 2

		if n == line1Index || n == line2Index || n == line3Index {
			lines = append(lines, line)

			if n == line3Index {
				char := ""

				line1 := lines[0]
				line2 := lines[1]
				line3 := lines[2]

				if char == "" {
					for j := 0; j < len(line1); j++ {
						char1 := string(line1[j])

						for k := 0; k < len(line2); k++ {
							char2 := string(line2[k])

							for l := 0; l < len(line3); l++ {
								char3 := string(line3[l])

								if char1 == char2 && char2 == char3 && char1 == char3 {
									char = char1
								}
							}
						}
					}
				}

				if char != "" {
					if unicode.IsUpper([]rune(char)[0]) {
						total += capitalLetters[char]
					} else {
						total += simpleLetters[char]
					}
				}

				lines = nil
				i++
			}
		}

		n++
	}

	return total
}
