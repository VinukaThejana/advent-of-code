package main

import "strings"

func getNewTotal() {
	scanner := readFile()

	for scanner.Scan() {
		line := scanner.Text()

		// Get the data
		data := strings.Split(line, " ")
		opponent := data[0]
		you := data[1]

		switch opponent {
		case string(A):
			if you == string(X) {
				total += getMatchScores(L)
				total += getObjectScores(S)
				break
			}

			if you == string(Y) {
				total += getMatchScores(D)
				total += getObjectScores(R)
				break
			}

			total += getMatchScores(W)
			total += getObjectScores(P)
			break

		case string(B):
			if you == string(X) {
				total += getMatchScores(L)
				total += getObjectScores(R)
				break
			}

			if you == string(Y) {
				total += getMatchScores(D)
				total += getObjectScores(P)
				break
			}

			total += getMatchScores(W)
			total += getObjectScores(S)
			break

		case string(C):
			if you == string(X) {
				total += getMatchScores(L)
				total += getObjectScores(P)
				break
			}

			if you == string(Y) {
				total += getMatchScores(D)
				total += getObjectScores(S)
				break
			}

			total += getMatchScores(W)
			total += getObjectScores(R)
			break
		}

	}
}
