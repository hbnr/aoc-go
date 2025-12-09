package main

import "strconv"

func SolvePartTwo(lines []string) int {
	totalInvalidIDs := 0
	spans := parseLine(lines[0])

	for _, span := range spans {
		for i := span[0]; i <= span[1]; i++ {
			if !validId2(i) {
				totalInvalidIDs += i
			}
		}
	}

	return totalInvalidIDs
}

func validId2(id int) bool {
	idText := strconv.Itoa(id)
	textLength := len(idText)

	for i := 1; i < textLength; i++ {
		if textLength%i != 0 {
			continue
		}
		pattern := idText[0:i]
		patternValid := false

		for n := 0; n < textLength; n += i {
			slice := idText[n : n+i]
			if slice != pattern {
				patternValid = true
				break
			}
		}
		if !patternValid {
			return false
		}
	}

	return true
}
