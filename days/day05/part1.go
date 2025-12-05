// Package day05 contains solutions for the Day 5 problems
package day05

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const rangeCount int64 = 190

type Range struct {
	Start int64
	End   int64
}

func isFresh(ranges []Range, id int64) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}

	return false
}

func Part1() int64 {
	file, err := os.Open("days/day05/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	ranges := make([]Range, 0, rangeCount)

	scanner := bufio.NewScanner(file)

	var lineIndex int64 = 0
	var solution int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		if lineIndex < rangeCount {
			parts := strings.Split(line, "-")
			start, startErr := strconv.ParseInt(parts[0], 10, 64)
			if startErr != nil {
				panic("Failed to parse start of range: " + startErr.Error())
			}

			end, endErr := strconv.ParseInt(parts[1], 10, 64)
			if endErr != nil {
				panic("Failed to parse end of range: " + endErr.Error())
			}

			ranges = append(ranges, Range{Start: start, End: end})
		}

		if lineIndex > rangeCount {
			id, idErr := strconv.ParseInt(line, 10, 64)
			if idErr != nil {
				panic("Failed to parse ID: " + idErr.Error())
			}

			if isFresh(ranges, id) {
				solution++
			}
		}

		lineIndex++
	}

	return solution
}
