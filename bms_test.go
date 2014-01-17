package bms

import (
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func TestBms1(t *testing.T) {
	haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
	needle := "bokko"

	assert(t, Search(haystack, needle) == 2)
}

func TestBms2(t *testing.T) {
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
