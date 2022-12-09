package main

import "fmt"

func dayOne() {
  scanner := readFile()
	scanner.Scan()

	line := scanner.Text()

	var items []string

	for i, c := range line {
		char := string(c)

		if i < 3 {
			items = append(items, char)
			continue
		}

		items = append(items, char)

		var buf []string
		seen := make(map[string]bool)

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				buf = append(buf, item)
			}
		}

		if len(buf) == 4 {
			fmt.Println("Found it!", char, i+1)
			break
		}

		items = append(items[:0], items[1:]...)
	}
}
