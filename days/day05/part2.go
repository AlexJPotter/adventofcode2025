// Package day05 contains solutions for the Day 5 problems
package day05

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func rangesOverlap(r1 Range, r2 Range) bool {
	if r1.Start > r2.End || r1.End < r2.Start {
		return false
	}

	return true
}

func mergeRanges(r1 Range, r2 Range) Range {
	start := min(r1.Start, r2.Start)
	end := max(r1.End, r2.End)
	return Range{Start: start, End: end}
}

func compactRanges(ranges []Range) []Range {
	for rightIndex := len(ranges) - 1; rightIndex > 0; rightIndex-- {
		for leftIndex := 0; leftIndex < rightIndex; leftIndex++ {
			leftRange := ranges[leftIndex]
			rightRange := ranges[rightIndex]

			if rangesOverlap(leftRange, rightRange) {
				ranges[leftIndex] = mergeRanges(leftRange, rightRange)
				ranges = append(ranges[:rightIndex], ranges[rightIndex+1:]...)
				break
			}
		}
	}

	return ranges
}

func Part2() int64 {
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

		if lineIndex >= rangeCount {
			break
		}

		lineIndex++
	}

	ranges = compactRanges(ranges)

	for _, r := range ranges {
		solution += (r.End - r.Start) + 1
	}

	return solution
}
