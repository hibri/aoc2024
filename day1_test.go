package aoc2024

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func TestCanGetLeftAndRightLists(t *testing.T) {

	input := []string{"1 3"}
	left, right, _ := SortLocations(input)
	firstLeft := left[0]
	firstRight := right[0]

	then.AssertThat(t, firstLeft, is.EqualTo(1))
	then.AssertThat(t, firstRight, is.EqualTo(3))
}

func TestCanHandleTwoLinesAsInput(t *testing.T) {

	input := []string{"1 3", "2 4"}
	left, right, _ := SortLocations(input)
	firstLeft := left[0]
	firstRight := right[0]

	then.AssertThat(t, firstLeft, is.EqualTo(1))
	then.AssertThat(t, firstRight, is.EqualTo(3))

}

func TestDistanceListAreSortedAscending(t *testing.T) {
	input := []string{"2 4", "1 3"}
	left, right, _ := SortLocations(input)
	firstLeft := left[0]
	firstRight := right[0]

	then.AssertThat(t, firstLeft, is.EqualTo(1))
	then.AssertThat(t, firstRight, is.EqualTo(3))
}
func TestReturnsDistancesForASimpleList(t *testing.T) {
	input := []string{"2 4", "1 3"}
	_, _, distance := SortLocations(input)

	then.AssertThat(t, distance[0], is.EqualTo(2))
	then.AssertThat(t, distance[1], is.EqualTo(2))
}
func TestReturnsDistancesForALongerList(t *testing.T) {
	input := []string{"3 4", "4 3", "2 5", "1 3", "3 9", "3 3"}
	_, _, distance := SortLocations(input)

	then.AssertThat(t, distance, is.EqualTo([]int{2, 1, 0, 1, 2, 5}))

}

func TestReturnsTotalDistanceForALongerList(t *testing.T) {

	exptectedTotalDistance := 11
	input := []string{"3 4", "4 3", "2 5", "1 3", "3 9", "3 3"}
	totalDistance := CalculateTotalDistance(input)
	then.AssertThat(t, totalDistance, is.EqualTo(exptectedTotalDistance))
}
