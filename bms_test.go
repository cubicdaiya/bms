package bms

import (
	"testing"
)

func TestSearch(t *testing.T) {

	tests := []struct {
		haystack string
		needle   string
		result   int
	}{
		{
			// SearchSuccess
			"bokkobokkkobokkkkobokkobokkkobokkkko",
			"bokko",
			2,
		},
		{
			// SearchFail
			"bokkobokkkobokkkkobokkobokkkobokkkko",
			"bokkkkko",
			0,
		},
		{
			// SearchMultibyte
			"bokkobokkko久保bokkkkobokko久保bokkkobokkkko",
			"久保",
			2,
		},
		{
			// EmptyHaystack
			"",
			"bokko",
			0,
		},
		{
			// EmptyNeedle
			"bokko",
			"",
			0,
		},
		{
			// ShorterHaystackThanNeedle
			"bokko",
			"bokkko",
			0,
		},
		{
			// SameHaystackAndNeedle
			"bokko",
			"bokko",
			1,
		},
		{
			// SameLengthHaystackAndNeedle
			"okkob",
			"bokko",
			0,
		},
	}

	for _, test := range tests {
		result := Search(test.haystack, test.needle)
		if result != test.result {
			t.Errorf("haystack: %s, needle: %s, want: %d, got: %d\n", test.haystack, test.needle, test.result, result)
		}
	}
}

func TestSearchBySkipTable(t *testing.T) {
	haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
	needle := "bokko"
	expected := 2

	table := BuildSkipTable(needle)
	got := SearchBySkipTable(haystack, needle, table)

	if got != expected {
		t.Errorf("haystack: %s, needle: %s, want: %d, got: %d\n", haystack, needle, expected, got)
	}
}

func TestSearchByInvalidSkipTable(t *testing.T) {
	haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
	needle := "bokko"
	needle2 := "cubicdaiya"
	expected := 1

	table := BuildSkipTable(needle2)
	got := SearchBySkipTable(haystack, needle, table)

	if got != expected {
		t.Errorf("haystack: %s, needle: %s, want: %d, got: %d\n", haystack, needle, expected, got)
	}
}
