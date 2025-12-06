// Package day06 contains solutions for the Day 6 problems
package day06

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part2() int64 {
	columns := parseColumns()

	columnWidths := make([]int, 0, len(columns))

	for _, column := range columns {
		columnWidths = append(columnWidths, getWidth(column))
	}

	columns = parseColumnsWidthAware(columns, columnWidths)

	var solution int64 = 0

	for _, column := range columns {
		solution += getColumnResult2(column)
	}

	return solution
}

func parseColumnsWidthAware(columns [][]string, columnWidths []int) [][]string {
	file, err := os.Open("days/day06/input.txt")
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		cursor := 0

		for columnIndex, columnWidth := range columnWidths {
			if cursor >= len(line) {
				break
			}

			endIndex := min(len(line), cursor+columnWidth)
			stringValue := line[cursor:endIndex]
			columns[columnIndex][lineIndex] = stringValue

			cursor += columnWidth + 1 // +1 to account for the space separator
		}

		lineIndex++
	}

	return columns
}

func getColumnResult2(column []string) int64 {
	width := getWidth(column)
	operator := strings.TrimSpace(column[len(column)-1])

	var result int64

	switch operator {
	case "+":
		result = 0
	case "*":
		result = 1
	default:
		panic("Unrecognised operator: " + operator)
	}

	for innerColumnIndex := range width {
		valueAtInnerColumn := getValueAtInnerColumn(column, innerColumnIndex)

		switch operator {
		case "+":
			result += valueAtInnerColumn
		case "*":
			result *= valueAtInnerColumn
		default:
			panic("Unrecognised operator: " + operator)
		}
	}

	return result
}

func getValueAtInnerColumn(column []string, innerColumnIndex int) int64 {
	height := getHeight(column)

	var result int64 = 0

	for rowIndex, stringValue := range column {
		if rowIndex == len(column)-1 {
			return result
		}

		if innerColumnIndex >= len(stringValue) {
			continue
		}

		if stringValue[innerColumnIndex] == ' ' {
			result = result / 10
			continue
		}

		exponent := height - (rowIndex + 1)
		digit, err := strconv.ParseInt(stringValue[innerColumnIndex:innerColumnIndex+1], 10, 64)
		if err != nil {
			panic("Failed to parse digit at inner column index: " + err.Error())
		}

		var multiplier int64 = 1

		for range exponent {
			multiplier *= 10
		}

		result += multiplier * digit
	}

	return result
}

func getHeight(column []string) int {
	return len(column) - 1 // Omit the operator
}

func getWidth(column []string) int {
	result := 0

	for index, value := range column {
		if index < len(column)-1 {
			result = max(result, len(value))
		}
	}

	return result
}
