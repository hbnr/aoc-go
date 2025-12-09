package main

import (
	"fmt"
	"time"

	"github.com/hbnr/aoc-go/internal/input"
	"github.com/hbnr/aoc-go/internal/stats"
)

const (
	sample = true
	year   = "2025"
	day    = "1"
)

func main() {
	start := time.Now()
	defer stats.LogTimeExecution(start)

	fmt.Printf("# Day %v", day)
	if sample {
		fmt.Println(" - Sample data")
	} else {
		fmt.Println()
	}

	fileName := input.GetFileName(year, day, sample)
	lines := input.MustReadLines(fileName)

	const startPoint = 50
	const limit = 100

	fmt.Println("Solution part one: ", SolvePartOne(lines, startPoint, limit))
	fmt.Println("Solution part two: ", SolvePartTwo(lines, startPoint, limit))
}
