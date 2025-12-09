package main

func SolvePartTwo(lines []string, startPoint, limit int) int {
	timesItHitZero := 0
	previousPosition := startPoint

	for _, line := range lines {
		timesItHitZero += passedZero(previousPosition, line, limit)
		previousPosition = dial(previousPosition, line, limit)
	}

	return timesItHitZero
}

func passedZero(position int, line string, limit int) int {
	countTimesHitZero := 0
	direction := getDirection(line)
	offset := getOffset(line)
	distanceToZero := 0

	if position != 0 {
		if direction == 1 {
			distanceToZero = limit - position
		} else {
			distanceToZero = position
		}
	}

	if distanceToZero == 0 && offset > limit {
		countTimesHitZero += offset / limit
	} else {
		if distanceToZero > 0 && distanceToZero <= offset {
			countTimesHitZero++
			rest := offset - distanceToZero
			countTimesHitZero += rest / limit
		}
	}

	return countTimesHitZero
}
