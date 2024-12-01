package aoc2024

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func TestCanGetLeftAndRightLists(t *testing.T) {

	input := []string{"1 3"}
	left, right := sortLocations(input)
	firstLeft := left[0]
	firstRight := right[0]

	then.AssertThat(t, firstLeft, is.EqualTo(1))
	then.AssertThat(t, firstRight, is.EqualTo(3))
}

func TestCanHandleTwoLinesAsInput(t *testing.T) {

	input := []string{"1 3", "2 4"}
	left, right := sortLocations(input)
	firstLeft := left[0]
	firstRight := right[0]

	then.AssertThat(t, firstLeft, is.EqualTo(1))
	then.AssertThat(t, firstRight, is.EqualTo(3))

}

// func TestDistanceListAreSortedAscending(t *testing.T) {
//
//		left, right := sortLocations()
//		firstLeft := left[0]
//		firstRight := right[0]
//
//		then.AssertThat(t, firstLeft, is.EqualTo(1))
//		then.AssertThat(t, firstRight, is.EqualTo(3))
//	}
func sortLocations(input []string) ([]int, []int) {
	var left []int
	var right []int
	for index, value := range input {
		var splitString = strings.Split(value, " ")
		left = append(left, 0)
		right = append(right, 0)
		left[index], _ = strconv.Atoi(splitString[0])
		right[index], _ = strconv.Atoi(splitString[1])
	}
	return left, right
}
