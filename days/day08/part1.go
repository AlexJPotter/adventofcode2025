// Package day08 contains solutions for the Day 8 problems
package day08

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type vec3 struct{ X, Y, Z float64 }

type pairwiseDistance struct {
	P, Q     vec3
	Distance float64
}

func Part1() int {
	boxLocations := parseInput()

	pairwiseDistances := make([]pairwiseDistance, 0)

	circuitIDsByLocation := make(map[vec3]int)
	circuitID := 0

	locationsByCircuitID := make(map[int]([]vec3))

	circuitSizesByID := make(map[int]int)

	for outerIndex, outerLocation := range boxLocations {
		for innerIndex := outerIndex + 1; innerIndex < len(boxLocations); innerIndex++ {
			innerLocation := boxLocations[innerIndex]

			p, q := ordered(outerLocation, innerLocation)

			distance := calculateDistance(p, q)

			pairwiseDistance := pairwiseDistance{P: p, Q: q, Distance: distance}
			pairwiseDistances = append(pairwiseDistances, pairwiseDistance)

			circuitIDsByLocation[p] = circuitID
			locationsByCircuitID[circuitID] = []vec3{p}
			circuitSizesByID[circuitID] = 1
			circuitID++

			circuitIDsByLocation[q] = circuitID
			locationsByCircuitID[circuitID] = []vec3{q}
			circuitSizesByID[circuitID] = 1
			circuitID++
		}
	}

	sort.Slice(pairwiseDistances, func(i, j int) bool {
		return pairwiseDistances[i].Distance < pairwiseDistances[j].Distance
	})

	connectionsToMake := 1000

	for index := range connectionsToMake {
		pairwiseDistance := pairwiseDistances[index]
		p := pairwiseDistance.P
		q := pairwiseDistance.Q

		pCircuitID := circuitIDsByLocation[p]
		qCircuitID := circuitIDsByLocation[q]

		if pCircuitID == qCircuitID {
			continue
		}

		for _, location := range locationsByCircuitID[qCircuitID] {
			locationsByCircuitID[pCircuitID] = append(locationsByCircuitID[pCircuitID], location)
			circuitIDsByLocation[location] = pCircuitID
		}

		circuitSizesByID[pCircuitID] += circuitSizesByID[qCircuitID]

		delete(locationsByCircuitID, qCircuitID)
		delete(circuitSizesByID, qCircuitID)
	}

	circuitSizes := make([]int, 0)

	for _, circuitSize := range circuitSizesByID {
		circuitSizes = append(circuitSizes, circuitSize)
	}

	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]
	})

	solution := 1

	circuitSizesToMultiply := 3

	for index := range circuitSizesToMultiply {
		solution *= circuitSizes[index]
	}

	return solution
}

func parseInput() []vec3 {
	file, err := os.Open("days/day08/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	boxLocations := make([]vec3, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		z, _ := strconv.ParseFloat(parts[2], 64)
		boxLocation := vec3{X: x, Y: y, Z: z}
		boxLocations = append(boxLocations, boxLocation)
	}

	return boxLocations
}

func calculateDistance(p, q vec3) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	dz := p.Z - q.Z
	return math.Sqrt((dx * dx) + (dy * dy) + (dz * dz))
}

// Puts the two points in a consistent order, which is useful for symmetric relationships
func ordered(p, q vec3) (vec3, vec3) {
	if p.X < q.X {
		return p, q
	}
	if q.X < p.X {
		return q, p
	}

	if p.Y < q.Y {
		return p, q
	}
	if q.Y < p.Y {
		return q, p
	}

	if p.Z < q.Z {
		return p, q
	}
	if q.Z < p.Z {
		return q, p
	}

	return p, q
}
