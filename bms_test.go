package bms

import (
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func TestSearchSuccess(t *testing.T) {
	haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
	needle := "bokko"

	assert(t, Search(haystack, needle) == 2)
}

func TestSearchFail(t *testing.T) {
	haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
	needle := "bokkkkko"

	assert(t, Search(haystack, needle) == 0)
}

func TestSearchMutibyte(t *testing.T) {
	haystack := "bokkobokkko久保bokkkkobokko久保bokkkobokkkko"
	needle := "久保"

	assert(t, Search(haystack, needle) == 2)
}

func TestEmtpyHaystack(t *testing.T) {
	haystack := ""
	needle := "bokko"

	assert(t, Search(haystack, needle) == 0)
}

func TestEmtpyNeedle(t *testing.T) {
	haystack := "bokko"
	needle := ""

	assert(t, Search(haystack, needle) == 0)
}

func TestShorterHaystackThanNeedle(t *testing.T) {
	haystack := "bokko"
	needle := "bokkko"

	assert(t, Search(haystack, needle) == 0)
}

func TestSameHaystackAndNeedle(t *testing.T) {
	haystack := "bokko"
	needle := "bokko"

	assert(t, Search(haystack, needle) == 1)
}


func TestSameLengthHaystackAndNeedle(t *testing.T) {
	haystack := "okkob"
	needle := "bokko"

	assert(t, Search(haystack, needle) == 0)
}

func TestSearchBySkipTable(t *testing.T) {
	haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
	needle := "bokko"

	table := BuildSkipTable(needle)
	assert(t, SearchBySkipTable(haystack, needle, table) == 2)
}

func TestSearchByInvalidSkipTable(t *testing.T) {
	haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
	needle := "bokko"
	needle2 := "cubicdaiya"

	table := BuildSkipTable(needle2)
	assert(t, SearchBySkipTable(haystack, needle, table) != 2)
}
