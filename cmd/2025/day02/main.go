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
	day    = "2"
)

// Sample
// 1: 1227775554
// 2: 4174379265
// Input
// 1: 24043483400
// 2: 38262920235

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
