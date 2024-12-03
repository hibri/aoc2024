package aoc2024

import (
	"bufio"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"io"
	"os"
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

func TestShouldBeUnSafeWhenLevelsDontIncreaseOrDecrease(t *testing.T) {

	reportLine := "8 6 4 4 1"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(false))
}

func TestShouldBeSafeWhenLevelsIncreased(t *testing.T) {

	reportLine := "1 3 6 7 9"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}

func TestShouldBeSafeWhenSecondLevelsRemoved(t *testing.T) {

	reportLine := "1 2 4 5"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}

func TestShouldBeSafeWhenThirdLevelIsRemoved(t *testing.T) {

	reportLine := "8 6 4 1"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}
func TestShouldBeUnSafeWhenAnyLevelIsRemoved(t *testing.T) {

	reportLine := "9 7 2 1"

	_, isSafe := ReadReport(reportLine)

	then.AssertThat(t, isSafe, is.EqualTo(false))
}

func TestReadDay2InputFromFile(t *testing.T) {

	file, _ := os.Open("./day2_input.txt")
	lineScanner := bufio.NewReader(file)
	safeReports := 0
	unsafeReports := 0
	for {
		line, err := lineScanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		_, safe := ReadReport(strings.TrimSpace(line))
		if safe {
			safeReports++
		} else {
			unsafeReports++
		}

	}
	then.AssertThat(t, safeReports, is.EqualTo(0))
	then.AssertThat(t, unsafeReports, is.EqualTo(0))
}

func ReadReport(input string) ([]int, bool) {

	var levels []int
	var changes []int
	var rawLevels = strings.Split(input, " ")
	var isSafe = false
	isIncrease := false
	isDecrease := false
	badLevels := 0

	for index, value := range rawLevels {
		levels = append(levels, 0)
		levels[index], _ = strconv.Atoi(value)
	}
	for index := 0; index < len(levels)-1; index++ {
		current := levels[index]
		next := levels[index+1]
		changes = append(changes, 0)
		diff := diff(next, current)
		if diff > 0 {
			isIncrease = true
		} else if diff < 0 {
			isDecrease = true
		} else if diff == 0 {
			isIncrease = false
			isDecrease = false
			badLevels++
		}

		changes[index] = diff
		if isSafeDiff(diff) {
			isSafe = true

		} else {
			isSafe = false
			badLevels++
			break

		}

	}
	isSafe = isSafe && ((isIncrease || isDecrease) && !(isIncrease && isDecrease) && !(!isDecrease && !isIncrease))

	return levels, isSafe
}


func isSafeDiff(diff int) bool {

	if diff < 0 {
		diff = diff * -1
	}
	return diff == 1 || diff == 2 || diff == 3
}

func diff(next int, current int) int {
	return next - current
}
