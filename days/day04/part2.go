// Package day04 contains solutions for the Day 4 problems
package day04

import (
	"bufio"
	"os"
)

func Part2() int64 {
	var grid Grid

	file, err := os.Open("days/day04/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var solution int64 = 0

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		for x, character := range line {
			location := Vec2{X: x, Y: y}
			setGridValue(&grid, location, character == '@')
		}

		y++
	}

	removedInLatestPass := 1 // Set this to >0 so we can enter the loop

	for removedInLatestPass > 0 {
		removedInLatestPass = 0

		for index, value := range grid {
			if !value {
				continue
			}

			location := getGridLocation(index)
			adjacentTowels := 0

			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}

					adjacentLocation := Vec2{X: location.X + x, Y: location.Y + y}

					if isValidGridLocation(adjacentLocation) && getGridValue(&grid, adjacentLocation) {
						adjacentTowels++
					}
				}
			}

			if adjacentTowels < 4 {
				solution++
				removedInLatestPass++
				setGridValue(&grid, location, false)
			}
		}
	}

	return solution
}
