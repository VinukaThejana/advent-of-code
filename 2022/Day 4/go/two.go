package main

func partialIntersections() int {
	scanner := readFile()
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		limit1, limit2 := extractor(line)

		// Check wether limit1 is within bounds of limit2
		if (limit1.Lower >= limit2.Lower && limit1.Lower <= limit2.Upper) || (limit1.Upper >= limit2.Lower && limit1.Upper <= limit2.Upper) {
			count++
			continue
		}

		// Check wether the limit2 is within the bounds of limit1
		if (limit2.Lower >= limit1.Lower && limit2.Lower <= limit1.Upper) || (limit2.Upper >= limit1.Lower && limit2.Upper <= limit1.Upper) {
			count++
			continue
		}

	}

	return count
}
