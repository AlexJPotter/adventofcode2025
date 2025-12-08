// Package day08 contains solutions for the Day 8 problems
package day08

import (
	"sort"
)

func Part2() float64 {
	boxLocations := parseInput()

	pairwiseDistances := make([]pairwiseDistance, 0)

	circuitIDsByLocation := make(map[vec3]int)
	circuitID := 0

	locationsByCircuitID := make(map[int]([]vec3))

	for outerIndex, outerLocation := range boxLocations {
		for innerIndex := outerIndex + 1; innerIndex < len(boxLocations); innerIndex++ {
			innerLocation := boxLocations[innerIndex]

			p, q := ordered(outerLocation, innerLocation)

			distance := calculateDistance(p, q)

			pairwiseDistance := pairwiseDistance{P: p, Q: q, Distance: distance}
			pairwiseDistances = append(pairwiseDistances, pairwiseDistance)

			circuitIDsByLocation[p] = circuitID
			locationsByCircuitID[circuitID] = []vec3{p}
			circuitID++

			circuitIDsByLocation[q] = circuitID
			locationsByCircuitID[circuitID] = []vec3{q}
			circuitID++
		}
	}

	sort.Slice(pairwiseDistances, func(i, j int) bool {
		return pairwiseDistances[i].Distance < pairwiseDistances[j].Distance
	})

	circuitCount := len(boxLocations)

	for index := range len(pairwiseDistances) {
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

		delete(locationsByCircuitID, qCircuitID)

		circuitCount--

		if circuitCount == 1 {
			return p.X * q.X
		}
	}

	return 0.0
}
