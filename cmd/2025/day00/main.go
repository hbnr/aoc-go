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
	day    = "0"
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
	fmt.Println("Solution part one: ", SolvePartOne(lines))
	fmt.Println("Solution part two: ", SolvePartTwo(lines))
}
