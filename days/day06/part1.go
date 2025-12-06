// Package day06 contains solutions for the Day 6 problems
package day06

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1() int64 {
	columns := parseColumns()

	var solution int64 = 0

	for _, column := range columns {
		solution += getColumnResult(column)
	}

	return solution
}

func parseColumns() [][]string {
	file, err := os.Open("days/day06/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	columns := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		for index, value := range parts {
			if index >= len(columns) {
				columns = append(columns, make([]string, 0))
			}

			columns[index] = append(columns[index], value)
		}
	}

	return columns
}

func getColumnResult(column []string) int64 {
	operator := column[len(column)-1]

	switch operator {
	case "+":
		return sumColumn(column)
	case "*":
		return multiplyColumn(column)
	default:
		panic("Unrecognised operator: " + operator)
	}
}

func sumColumn(column []string) int64 {
	var result int64 = 0

	for _, stringValue := range column {
		value, err := strconv.ParseInt(stringValue, 10, 64)

		if err == nil {
			result += value
		}
	}

	return result
}

func multiplyColumn(column []string) int64 {
	var result int64 = 1

	for _, stringValue := range column {
		value, err := strconv.ParseInt(stringValue, 10, 64)

		if err == nil {
			result *= value
		}
	}

	return result
}
