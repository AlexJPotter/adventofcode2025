// Package day03 contains solutions for the Day 1 problems
package day03

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func zeroFromIndex(array *[12]int64, index int) {
	for i := index; i < len(array); i++ {
		array[i] = 0
	}
}

func Part2() int64 {
	file, err := os.Open("days/day03/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var solution int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		var digits [12]int64
		zeroFromIndex(&digits, 0)

		for charIndex := range len(line) {
			character := line[charIndex:(charIndex + 1)]

			currentDigit, err := strconv.ParseInt(character, 10, 64)
			if err != nil {
				panic("Failed to parse digit: " + err.Error())
			}

			done := false

			for digitIndex := range len(digits) {
				digitsLeftToFillAfterThis := len(digits) - (digitIndex + 1)
				charactersLeftAfterThis := len(line) - (charIndex + 1)

				if (digitsLeftToFillAfterThis <= charactersLeftAfterThis) && (currentDigit > digits[digitIndex]) {
					digits[digitIndex] = currentDigit
					zeroFromIndex(&digits, digitIndex+1)
					done = true
				}

				if done {
					break
				}
			}
		}

		var joltage int64 = 0

		for digitIndex := range len(digits) {
			digit := digits[digitIndex]
			exponent := len(digits) - (digitIndex + 1)
			joltage += int64(math.Pow10(exponent)) * digit
		}

		solution += joltage
	}

	return solution
}
