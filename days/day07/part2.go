// Package day07 contains solutions for the Day 7 problems
package day07

func Part2() int {
	entryPoint, splitterLocations, y := parseInput()

	beams := map[vec2]int{entryPoint: 1}

	for range y - 1 {
		updatedBeams := processBeams2(beams, splitterLocations)
		beams = updatedBeams
	}

	solution := 0

	for _, count := range beams {
		solution += count
	}

	return solution
}

func processBeams2(beams map[vec2]int, splitterLocations map[vec2]bool) map[vec2]int {
	newBeams := make(map[vec2]int, 0)

	for coordinate, count := range beams {
		coordinateBelow := vec2{X: coordinate.X, Y: coordinate.Y + 1}

		_, hasHitSplitter := splitterLocations[coordinateBelow]

		if hasHitSplitter {
			left := vec2{X: coordinateBelow.X - 1, Y: coordinateBelow.Y}
			increment(newBeams, left, count)

			right := vec2{X: coordinateBelow.X + 1, Y: coordinateBelow.Y}
			increment(newBeams, right, count)
		} else {
			increment(newBeams, coordinateBelow, count)
		}
	}

	return newBeams
}

func increment(beams map[vec2]int, coordinate vec2, count int) {
	currentCount, hasCurrentCount := beams[coordinate]

	if hasCurrentCount {
		beams[coordinate] = currentCount + count
	} else {
		beams[coordinate] = count
	}
}
