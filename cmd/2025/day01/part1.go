package main

import (
	"strconv"
)

func SolvePartOne(lines []string, startPoint, limit int) int {
	currentPosition := startPoint
	timesItHitZero := 0

	for _, line := range lines {
		currentPosition = dial(currentPosition, line, limit)
		if currentPosition == 0 {
			timesItHitZero++
		}
	}

	return timesItHitZero
}

func dial(startPosition int, line string, limit int) int {
	if len(line) < 1 {
		return startPosition
	}

	direction := getDirection(line)
	offset := getOffset(line) % limit
	newPosition := (startPosition + (direction * offset)) % limit

	if newPosition < 0 {
		return newPosition + limit
	}

	return newPosition
}

func getDirection(line string) int {
	if line[0] == 'L' {
		return -1
	}

	return 1
}

func getOffset(line string) int {
	if len(line) < 1 {
		return 0
	}

	offset := line[1:]
	num, err := strconv.Atoi(offset)
	if err != nil {
		return 0
	}

	return num
}
