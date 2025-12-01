// Package day01 contains solutions for the Day 1 problems
package day01

import (
	"bufio"
	"os"
	"strconv"
)

func Part1() int {
	file, err := os.Open("days/day01/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	password := 0
	var position int64 = 50

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0:1]
		turns, err := strconv.ParseInt(line[1:], 10, 16)
		if err != nil {
			panic("Failed to parse number of turns: " + err.Error())
		}

		switch direction {
		case "L":
			{
				// The modulo operator can return negative numbers
				position = (((position - turns) % 100) + 100) % 100
				break
			}
		case "R":
			{
				position = (position + turns) % 100
				break
			}
		}

		if position == 0 {
			password++
		}
	}

	return password
}
