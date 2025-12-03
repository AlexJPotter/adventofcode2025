// Package day03 contains solutions for the Day 1 problems
package day03

import (
	"bufio"
	"os"
	"strconv"
)

func Part1() int64 {
	file, err := os.Open("days/day03/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var solution int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		var left int64 = 0
		var right int64 = 0

		for charIndex := range len(line) {
			character := line[charIndex:(charIndex + 1)]

			digit, err := strconv.ParseInt(character, 10, 64)
			if err != nil {
				panic("Failed to parse digit: " + err.Error())
			}

			if digit > left && charIndex < len(line)-1 {
				left = digit
				right = 0
			} else if digit > right {
				right = digit
			}
		}

		joltage := (10 * left) + right
		solution += joltage
	}

	return solution
}
