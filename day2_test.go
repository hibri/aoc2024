package aoc2024

import (
	"bufio"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"io"
	"os"
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

	reportLine := "1 3 2 4 5"

	report, isSafe := ReadReport(reportLine)

	_, isSafe = ReCheckReport(report)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}

func TestShouldBeSafeWhenThirdLevelIsRemoved(t *testing.T) {

	reportLine := "8 6 4 4 1"

	report, isSafe := ReadReport(reportLine)

	_, isSafe = ReCheckReport(report)

	then.AssertThat(t, isSafe, is.EqualTo(true))
}
func TestShouldBeSafeWhenAnyLevelIsRemoved(t *testing.T) {

	reportLine := "15 18 20 21 23 25 28 32"

	report, isSafe := ReadReport(reportLine)
	_, isSafe = ReCheckReport(report)

	then.AssertThat(t, isSafe, is.EqualTo(true))
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
	_ = file.Close()
	then.AssertThat(t, safeReports, is.EqualTo(510))
	then.AssertThat(t, unsafeReports, is.EqualTo(490))
}

func TestReadDay2InputWithDampningFromFile(t *testing.T) {

	file, _ := os.Open("./day2_input.txt")
	lineScanner := bufio.NewReader(file)
	safeReports := 0
	unsafeReports := 0
	for {
		line, err := lineScanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		report, safe := ReadReport(strings.TrimSpace(line))
		if safe {
			safeReports++
		} else {
			_, safe := ReCheckReport(report)
			if safe {
				safeReports++
			}
			unsafeReports++

		}

	}
	_ = file.Close()
	then.AssertThat(t, safeReports, is.EqualTo(553))
	then.AssertThat(t, unsafeReports, is.EqualTo(490))
}
