// Package day07 contains solutions for the Day 7 problems
package day07

import (
	"bufio"
	"os"
)

const (
	entryPointChar rune = 'S'
	splitterChar   rune = '^'
)

type vec2 struct{ X, Y int }

func Part1() int {
	entryPoint, splitterLocations, y := parseInput()

	beams := map[vec2]bool{entryPoint: true}
	splitCount := 0

	for range y {
		updatedBeams, splitsPerformed := processBeams(beams, splitterLocations)
		beams = updatedBeams
		splitCount += splitsPerformed
	}

	return splitCount
}

func parseInput() (vec2, map[vec2]bool, int) {
	file, err := os.Open("days/day07/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entryPoint vec2
	splitterLocations := make(map[vec2]bool)

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		for x, character := range line {
			coordinate := vec2{X: x, Y: y}

			switch character {
			case entryPointChar:
				entryPoint = coordinate
			case splitterChar:
				splitterLocations[coordinate] = true
			}
		}

		y++
	}

	return entryPoint, splitterLocations, y
}

func processBeams(beams map[vec2]bool, splitterLocations map[vec2]bool) (map[vec2]bool, int) {
	newBeams := make(map[vec2]bool)
	splitsPerformed := 0

	for coordinate := range beams {
		coordinateBelow := vec2{X: coordinate.X, Y: coordinate.Y + 1}

		_, hasHitSplitter := splitterLocations[coordinateBelow]

		if hasHitSplitter {
			left := vec2{X: coordinateBelow.X - 1, Y: coordinateBelow.Y}
			right := vec2{X: coordinateBelow.X + 1, Y: coordinateBelow.Y}
			newBeams[left] = true
			newBeams[right] = true
			splitsPerformed++
		} else {
			newBeams[coordinateBelow] = true
		}
	}

	return newBeams, splitsPerformed
}
