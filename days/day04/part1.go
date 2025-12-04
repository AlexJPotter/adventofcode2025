// Package day04 contains solutions for the Day 4 problems
package day04

import (
	"bufio"
	"os"
)

const (
	gridWidth  = 140
	gridHeight = gridWidth
	gridSize   = gridWidth * gridHeight
)

type Vec2 struct {
	X int
	Y int
}

type Grid = [gridSize]bool

func getGridIndex(location Vec2) int {
	return (location.Y * gridWidth) + location.X
}

func getGridLocation(index int) Vec2 {
	return Vec2{X: index % gridWidth, Y: index / gridWidth}
}

func isValidGridLocation(location Vec2) bool {
	return location.X >= 0 && location.X < gridWidth && location.Y >= 0 && location.Y < gridHeight
}

func getGridValue(grid *Grid, location Vec2) bool {
	return grid[getGridIndex(location)]
}

func setGridValue(grid *Grid, location Vec2, value bool) {
	grid[getGridIndex(location)] = value
}

func Part1() int64 {
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
		}
	}

	return solution
}
