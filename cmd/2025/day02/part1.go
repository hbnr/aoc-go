package main

import (
	"log"
	"strconv"
	"strings"
)

func SolvePartOne(lines []string) int {
	totalInvalidIDs := 0
	spans := parseLine(lines[0])

	for _, span := range spans {
		for i := span[0]; i <= span[1]; i++ {
			if !validId(i) {
				totalInvalidIDs += i
			}
		}
	}

	return totalInvalidIDs
}

type Span [2]int

func parseLine(line string) []Span {
	spans := []Span{}
	valueSpans := strings.Split(line, ",")

	for _, span := range valueSpans {
		spans = append(spans, parseSpan(span))
	}
	return spans
}

func parseSpan(line string) Span {
	split := strings.Split(line, "-")
	if len(split) != 2 {
		log.Fatal("Invalid line - parseSpan", split, line)
	}

	var span Span

	for i := 0; i < 2; i++ {
		span[i], _ = strconv.Atoi(split[i])
	}

	return span
}

func validId(id int) bool {
	idText := strconv.Itoa(id)
	textLength := len(idText)
	textLengthIsEven := textLength%2 == 0

	if !textLengthIsEven {
		return true
	}

	halfLength := textLength / 2
	first := idText[:halfLength]
	second := idText[halfLength:]

	return first != second
}
