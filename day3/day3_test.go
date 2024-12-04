package day3

import (
	"bufio"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var corruptedMemory = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestCanFindAMulSequenceInASimpleString(t *testing.T) {

	matches := getInstructions(corruptedMemory)

	then.AssertThat(t, len(matches), is.EqualTo(4))
}

func TestCanExtractSequencesInASimpleString(t *testing.T) {

	matches := getInstructions(corruptedMemory)

	then.AssertThat(t, len(matches), is.EqualTo(4))
	then.AssertThat(t, matches[0][0], is.EqualTo("mul(2,4)"))

}

func TestCanExtractNumbersFromASequencesInASimpleString(t *testing.T) {

	matches := getInstructions(corruptedMemory)
	expectedX, _ := strconv.Atoi(matches[0][1])
	expectedY, _ := strconv.Atoi(matches[0][2])

	then.AssertThat(t, expectedX, is.EqualTo(2))
	then.AssertThat(t, expectedY, is.EqualTo(4))
}

func TestCanMultiplyNumbersFromASequencesInASimpleString(t *testing.T) {

	matches := getInstructions(corruptedMemory)
	expectedResult := multiply(matches[0])

	then.AssertThat(t, expectedResult, is.EqualTo(8))
}

func TestShouldSumMultiplicationResults(t *testing.T) {

	expectedTotal := 161

	matches := getInstructions(corruptedMemory)
	total := executeInstructions(matches)
	then.AssertThat(t, total, is.EqualTo(expectedTotal))

}

func TestShouldSumResultsWithDay3Input(t *testing.T) {
	file, _ := os.Open("./day3_input.txt")
	lineScanner := bufio.NewReader(file)
	total := 0
	for {
		line, err := lineScanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		instructions := getInstructions(strings.TrimSpace(line))
		total = executeInstructions(instructions)

	}
	_ = file.Close()

	then.AssertThat(t, total, is.EqualTo(0))
}

func executeInstructions(matches [][]string) int {
	total := 0
	for _, instruction := range matches {
		total += multiply(instruction)
	}
	return total
}

func multiply(match []string) int {
	x := match[1]
	expectedX, _ := strconv.Atoi(x)
	y := match[2]
	expectedY, _ := strconv.Atoi(y)
	expectedResult := expectedX * expectedY
	return expectedResult
}

func getInstructions(corruptedMemory string) [][]string {
	regex := regexp.MustCompile(`(?:mul)\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(corruptedMemory, -1)
	return matches
}
