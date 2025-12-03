package aoc

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type CounterElf struct {
	Sum int
}

func byOccurances(value string, index int) bool {
	pattern := string(value[:index])
	var match bool
	for i := 0; i < len(value); i += index {
		if len(value) < i+index {
			match = false
			break
		}
		if value[i:i+index] == pattern {
			match = true
		} else {
			match = false
			break
		}
	}
	if match {
		return true
	}

	if index+1 < len(value)/2 {
		return byOccurances(value, index+1)
	}
	return false
}

func repeatingPattern(value string, dayone bool) bool {
	middle := len(value) / 2
	firsthalf := value[:middle]
	secondhalf := value[middle:]

	if dayone {
		return firsthalf == secondhalf
	}

	if firsthalf == secondhalf {
		return true
	}
	return byOccurances(value, 1)
}

func parseInput(r string) (lower int, upper int) {
	bounds := strings.Split(r, "-")

	if len(bounds) != 2 {
		panic("invalid range: " + r)
	}
	lower, err := strconv.Atoi(bounds[0])
	if err != nil {
		panic(err)
	}
	upper, err = strconv.Atoi(bounds[1])
	if err != nil {
		panic(err)
	}
	return
}

func DayTwo() {
	data, err := os.ReadFile("aoc/day_2_input.txt")
	if err != nil {
		panic(err)
	}

	elfDayOne := CounterElf{Sum: 0}
	elfDayTwo := CounterElf{Sum: 0}

	for r := range strings.SplitSeq(strings.TrimRight(string(data), "\n"), ",") {
		lower, upper := parseInput(r)

		for i := lower; i <= upper; i++ {
			value := strconv.Itoa(i)
			if len(value) == 1 {
				continue
			}

			if repeatingPattern(value, true) {
				elfDayOne.Sum += i
			}
			if repeatingPattern(value, false) {
				elfDayTwo.Sum += i
			}
		}
	}

	slog.Info("Part 1:", "total", elfDayOne.Sum)
	slog.Info("Part 2:", "total", elfDayTwo.Sum)
}
