package aoc2024

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"strconv"
	"strings"
	"testing"
)

func TestShouldReadLevelsFromReport(t *testing.T) {

	var reportLine = "7 6 4 2 1"
	expectedReportLine := []int{7, 6, 4, 2, 1}

	reportedLine, _ := ReadReport(reportLine)

	then.AssertThat(t, reportedLine, is.EqualTo(expectedReportLine))
}

func TestShouldBeSafeWhenLevelsDecreasingBy1(t *testing.T) {

	reportLine := "7 6"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}
func TestShouldBeSafeWhenLevelsDecreasingBy2(t *testing.T) {

	reportLine := "7 5"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}
func TestShouldBeSafeWhenLevelsDecreasingBy3(t *testing.T) {

	reportLine := "7 4"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}

func TestShouldBeSafeWhenLevelsDecreasingForAReportLine(t *testing.T) {

	reportLine := "7 6 4 2 1"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}

func TestShouldBeUnSafeWhenLevelsDecreasingForAReportLineByMoreThan3(t *testing.T) {

	reportLine := "7 4 4 0 1"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(false))
}

func TestShouldBeSafeWhenLevelsIncreaseBy1(t *testing.T) {

	reportLine := "1 2"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}

func TestShouldBeUnSafeWhenLevelsIncreaseBy5(t *testing.T) {

	reportLine := "1 2 7 8 9"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(false))
}

func TestShouldBeUnSafeWhenLevelsDecreaseBy4(t *testing.T) {

	reportLine := "9 7 6 2 1"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(false))
}

func TestShouldBeUnSafeWhenLevelsIncreaseAndDecrease(t *testing.T) {

	reportLine := "1 3 2 4 5"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(false))
}
func ReadReport(input string) ([]int, bool) {

	var levels []int
	var rawLevels = strings.Split(input, " ")
	var isSafe = false

	for index, value := range rawLevels {
		levels = append(levels, 0)
		levels[index], _ = strconv.Atoi(value)
	}
	for index := 0; index < len(levels)-1; index++ {
		current := levels[index]
		next := levels[index+1]
		if isSafeDiff(current, next) {
			isSafe = true
			safeDir = true

		} else {
			isSafe = false
			unSafeDir = true

		}

	}

	return levels, isSafe
}

func isSafeDiff(current int, next int) bool {

	diff := current - next
	if diff < 0 {
		diff = -diff
	}
	return diff == 1 || diff == 2 || diff == 3
}
