package main

func totalIntersection() int {
	scanner := readFile()
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		limit1, limit2 := extractor(line)

		// Check wether the limit1 fully contains limit2
		if (limit2.Lower >= limit1.Lower) && (limit2.Upper <= limit1.Upper) {
			count++
			continue
		}

		// Check wether the limit2 fully contains limit1
		if (limit1.Lower >= limit2.Lower) && (limit1.Upper <= limit2.Upper) {
			count++
			continue
		}
	}

	return count
}
