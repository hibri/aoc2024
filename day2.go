package aoc2024

import (
	"slices"
	"strconv"
	"strings"
)

func ReadReport(input string) ([]int, bool) {

	var levels []int
	var rawLevels = strings.Split(input, " ")
	var isSafe = false

	for index, value := range rawLevels {
		levels = append(levels, 0)
		levels[index], _ = strconv.Atoi(value)
	}
	changes := getChanges(levels)

	allChangesAreWithinSafeLimits := areAllChangesWithinSafeLimits(changes)
	//numberOfUnsafeLevels := countUnsafeLevels(changes)
	hasIncreasingChanges := hasIncreasingChanges(changes)
	decreasingChanges := hasDecreasingChanges(changes)

	isSafe = allChangesAreWithinSafeLimits && (hasIncreasingChanges || decreasingChanges) && !(hasIncreasingChanges && decreasingChanges)

	return levels, isSafe
}

func getChanges(levels []int) []int {
	var changes []int
	for index := 0; index < len(levels)-1; index++ {
		changes = append(changes, 0)

		changes[index] = diff(levels[index+1], levels[index])

	}
	return changes
}

func hasDecreasingChanges(changes []int) bool {
	hasDecreasingChanges := false

	for _, change := range changes {
		hasDecreasingChanges = change < 0
		if !hasDecreasingChanges {
			break
		}
	}

	return hasDecreasingChanges
}

func hasIncreasingChanges(changes []int) bool {
	hasIncreasingChanges := false

	for _, change := range changes {
		hasIncreasingChanges = change > 0
		if !hasIncreasingChanges {
			break
		}
	}
	return hasIncreasingChanges
}

func countUnsafeLevels(changes []int) int {
	countUnsafeLevels := 0
	for _, change := range changes {

		if !isSafeLevelChange(change) {
			countUnsafeLevels++
		}
	}
	return countUnsafeLevels
}

func areAllChangesWithinSafeLimits(changes []int) bool {
	var isSafe = false
	for _, change := range changes {

		isSafe = isSafeLevelChange(change)
		if !isSafe {
			break
		}
	}
	return isSafe
}
func isSafeLevelChange(diff int) bool {

	if diff < 0 {
		diff = diff * -1
	}
	return diff >= 1 && diff <= 3
}

func diff(next int, current int) int {
	return next - current
}

func ReCheckReport(report []int) ([]int, bool) {

	var isSafe bool
	for i := 0; i < len(report); i++ {
		var subset []int
		if i == 0 {
			subset = report[i+1 : len(report)]
		} else if i == len(report) {
			subset = report[0 : len(report)-1]
		} else {
			subset = slices.Concat(report[0:i], report[i+1:len(report)])
		}
		changes := getChanges(subset)
		allChangesAreWithinSafeLimits := areAllChangesWithinSafeLimits(changes)
		//numberOfUnsafeLevels := countUnsafeLevels(changes)
		hasIncreasingChanges := hasIncreasingChanges(changes)
		decreasingChanges := hasDecreasingChanges(changes)

		isSafe = allChangesAreWithinSafeLimits && (hasIncreasingChanges || decreasingChanges) && !(hasIncreasingChanges && decreasingChanges)
		if isSafe {
			break
		}
	}
	return report, isSafe
}
