// Package main is the main package for Advent of Code solutions.
package main

import (
	"log/slog"
	"mseppae/adventofcode/2025/aoc"
)

func main() {
	slog.Info("Day one: count dial zero hits")
	aoc.DayOne()
	slog.Info("----------------")
	slog.Info("Day two: count invalid passwords")
	aoc.DayTwo()
}
