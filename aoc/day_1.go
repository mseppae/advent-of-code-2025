// Package aoc is the package for Advent of Code solutions.
package aoc

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Dial struct {
	Position         int
	Maximum          int
	ZeroHits         int
	TurnOverZeroHits int
}

func (d *Dial) Turn(ticks string) {
	direction := ticks[0]
	amount := ticks[1:]
	a, err := strconv.Atoi(amount)
	if err != nil {
		panic(err)
	}
	turns := a / d.Maximum
	if turns > 0 {
		d.TurnOverZeroHits += turns
	}
	aMod := a % d.Maximum

	switch direction {
	case 'L':
		oldPosition := d.Position
		d.Position -= aMod
		if d.Position < 0 {
			if oldPosition != 0 {
				d.TurnOverZeroHits++
			}
			d.Position += d.Maximum
		}
	case 'R':
		d.Position += aMod
		if d.Position > d.Maximum {
			d.TurnOverZeroHits++
		}
		if d.Position >= d.Maximum {
			d.Position -= d.Maximum
		}
	}
	if d.Position == 0 {
		d.ZeroHits++
	}
}

func DayOne() {
	dial := Dial{Position: 50, Maximum: 100, ZeroHits: 0, TurnOverZeroHits: 0}
	data, err := os.ReadFile("aoc/day_1_input.txt")
	if err != nil {
		panic(err)
	}
	for instruction := range strings.SplitSeq(string(data), "\n") {
		if instruction == "" {
			continue
		}
		dial.Turn(instruction)
	}
	slog.Info("Part 1:", "total", dial.ZeroHits)
	slog.Info("Part 2:", "total", dial.ZeroHits+dial.TurnOverZeroHits)
}
