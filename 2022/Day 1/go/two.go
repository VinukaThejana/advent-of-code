package main

import (
	"sort"
	"strconv"
	"strings"
)

func getTheTotalCalsOfTheTopThreeElfs() int {
	scanner := readFile()

	var elfs []int
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elfs = append(elfs, total)
			total = 0
			continue
		}

		cal, err := strconv.Atoi(strings.Split(line, "\n")[0])
		if err != nil {
			panic(err)
		}

		total += cal

	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))
	total = elfs[0] + elfs[1] + elfs[2]

	return total
}
