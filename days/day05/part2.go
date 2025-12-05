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

func min(x int64, y int64) int64 {
	if x < y {
		return x
	}

	return y
}

func max(x int64, y int64) int64 {
	if x > y {
		return x
	}

	return y
}

func mergeRanges(r1 Range, r2 Range) Range {
	start := min(r1.Start, r2.Start)
	end := max(r1.End, r2.End)
	return Range{Start: start, End: end}
}

func dropAtIndex(ranges *[rangeCount]Range, rangesSize int, index int) {
	for i := index; i < rangesSize; i++ {
		ranges[i] = ranges[i+1]
	}
}

func compactRanges(ranges *[rangeCount]Range, rangesSize int) int {
	newRangesSize := rangesSize

	for rightIndex := rangesSize - 1; rightIndex > 0; rightIndex-- {
		for leftIndex := 0; leftIndex < rightIndex; leftIndex++ {
			leftRange := ranges[leftIndex]
			rightRange := ranges[rightIndex]

			if rangesOverlap(leftRange, rightRange) {
				ranges[leftIndex] = mergeRanges(leftRange, rightRange)
				newRangesSize -= 1
				dropAtIndex(ranges, newRangesSize, rightIndex)
				break
			}
		}
	}

	return newRangesSize
}

func Part2() int64 {
	file, err := os.Open("days/day05/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	var ranges [rangeCount]Range
	rangesSize := 0

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

			ranges[lineIndex] = Range{Start: start, End: end}
			rangesSize++
		}

		if lineIndex >= rangeCount {
			break
		}

		lineIndex++
	}

	rangesSize = compactRanges(&ranges, rangesSize)

	for index, r := range ranges {
		if index >= rangesSize {
			return solution
		}

		solution += (r.End - r.Start) + 1
	}

	return solution
}
