// Package day01 contains solutions for the Day 1 problems
package day01

import (
	"bufio"
	"os"
	"strconv"
)

func Part2() int64 {
	file, err := os.Open("days/day01/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	var password int64 = 0
	var position int64 = 50

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0:1]
		turns, err := strconv.ParseInt(line[1:], 10, 16)
		if err != nil {
			panic("Failed to parse number of turns: " + err.Error())
		}

		startedOnZero := position == 0

		switch direction {
		case "L":
			{
				position = position - turns
				break
			}
		case "R":
			{
				position = position + turns
				break
			}
		}

		crossings := position / 100

		// Using <= catches the case where we've gone "down" to 0 without crossing it
		if position <= 0 {
			crossings = -crossings

			// Account for the first instance of crossing 0 as we go down into negative numbers
			// (unless of course we started on 0, in which case this adjustment isn't needed)
			if !startedOnZero {
				crossings++
			}
		}

		password += crossings

		// Adjust position so that it definitely falls in 0..99
		position = ((position % 100) + 100) % 100
	}

	return password
}
