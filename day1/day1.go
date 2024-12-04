package day1

import (
	"slices"
	"strconv"
	"strings"
)

func SortLocations(input []string) ([]int, []int, []int) {
	var left []int
	var right []int
	var diffs []int
	for index, value := range input {
		var splitString = strings.Split(value, "   ")
		left = append(left, 0)
		right = append(right, 0)
		left[index], _ = strconv.Atoi(splitString[0])
		right[index], _ = strconv.Atoi(splitString[1])
	}
	slices.Sort(left)
	slices.Sort(right)
	for index, _ := range left {
		diff := left[index] - right[index]
		if diff < 0 {
			diff *= -1
		}
		diffs = append(diffs, diff)
	}

	return left, right, diffs
}

func CalculateTotalDistance(input []string) int {
	total := 0
	_, _, distances := SortLocations(input)
	for _, value := range distances {
		total += value
	}

	return total
}

func CalculateSimilarity(input []string) int {

	total := 0
	left, right, _ := SortLocations(input)
	var similarityScore []int
	for _, leftValue := range left {
		score := 0
		for _, rightValue := range right {
			if leftValue == rightValue {
				score++
			} else if rightValue > leftValue {
				break //break the innerloop, because lists are sorted, we don't need to go further
			}
		}
		similarityScore = append(similarityScore, leftValue*score)
	}
	for _, value := range similarityScore {
		total += value
	}
	return total
}
