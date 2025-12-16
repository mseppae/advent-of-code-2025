package aoc

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type JoltageBank struct {
	Total int
}

func (j *JoltageBank) Add(amount int) {
	j.Total = amount
}

func totalJoltageByIndices(bank string, indices []int) int {
	var total string
	for _, index := range indices {
		total += string(bank[index])
	}
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		panic(err)
	}
	return totalInt
}

func findLargest(bank string, leftindex int, remaining int) int {
	largest := 0
	index := -1

	for i := leftindex; i < len(bank)-remaining; i++ {
		voltage, err := strconv.Atoi(string(bank[i]))
		if err != nil {
			panic(err)
		}
		if voltage > largest {
			largest = voltage
			index = i
		}
	}

	return index
}

func joltagesForBatteryCount(data []byte, batterycount int) []*JoltageBank {
	var jBanks []*JoltageBank

	for bank := range strings.SplitSeq(string(data), "\n") {
		if bank == "" {
			continue
		}
		jBank := &JoltageBank{Total: 0}
		jBanks = append(jBanks, jBank)
		var indices []int
		index := 0
		remaining := batterycount - 1
		for range batterycount {
			index = findLargest(bank, index, remaining)
			indices = append(indices, index)
			index++
			remaining--
		}
		total := totalJoltageByIndices(bank, indices)
		jBank.Add(total)
	}
	return jBanks
}

func DayThree() {
	data, err := os.ReadFile("aoc/day_3_input.txt")
	if err != nil {
		panic(err)
	}

	jBanksPartOne := joltagesForBatteryCount(data, 2)
	jBanksPartTwo := joltagesForBatteryCount(data, 12)
	var totalPartOne int
	var totalPartTwo int
	for _, jBank := range jBanksPartOne {
		totalPartOne += jBank.Total
	}
	for _, jBank := range jBanksPartTwo {
		totalPartTwo += jBank.Total
	}

	slog.Info("Part 1:", "total", totalPartOne)
	slog.Info("Part 2:", "total", totalPartTwo)
}
