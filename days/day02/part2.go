// Package day02 contains solutions for the Day 2 problems
package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidWithSequenceLength(id string, sequenceLength int) bool {
	length := len(id)

	if length%sequenceLength != 0 {
		return false
	}

	repetitions := length / sequenceLength

	for sequenceIndex := range sequenceLength {
		expected := id[sequenceIndex]

		for repetitionIndex := range repetitions {
			if id[(repetitionIndex*sequenceLength)+sequenceIndex] != expected {
				return false
			}
		}
	}

	return true
}

func isInvalidID(id string) bool {
	length := len(id)

	if length < 2 {
		return false
	}

	for sequenceLength := 1; sequenceLength <= length/2; sequenceLength++ {
		if isInvalidWithSequenceLength(id, sequenceLength) {
			return true
		}
	}

	return false
}

func Part2() uint64 {
	file, err := os.Open("days/day02/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(scanCommaSeparated)

	var solution uint64 = 0

	for scanner.Scan() {
		id := scanner.Text()

		parts := strings.Split(strings.TrimSpace(id), "-")

		if len(parts) != 2 {
			panic(fmt.Sprint("Expected 2 parts but got ", len(parts)))
		}

		rangeStart, err := strconv.ParseUint(parts[0], 10, 64)
		if err != nil {
			panic("Failed to parse range start: " + err.Error())
		}

		rangeEnd, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			panic("Failed to parse range end: " + err.Error())
		}

		for id := rangeStart; id <= rangeEnd; id++ {
			stringID := strconv.FormatUint(id, 10)

			if isInvalidID(stringID) {
				solution += id
			}
		}
	}

	return solution
}
