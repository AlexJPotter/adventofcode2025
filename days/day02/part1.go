// Package day02 contains solutions for the Day 2 problems
package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanCommaSeparated(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := range len(data) {
		if data[i] == ',' {
			return i + 1, data[0:i], nil
		}
	}

	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func isValidID(id string) bool {
	length := len(id)

	if length%2 != 0 {
		return true
	}

	sequenceLength := length / 2

	for i := range sequenceLength {
		if id[i] != id[sequenceLength+i] {
			return true
		}
	}

	return false
}

func Part1() uint64 {
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

			if !isValidID(stringID) {
				solution += id
			}
		}
	}

	return solution
}
