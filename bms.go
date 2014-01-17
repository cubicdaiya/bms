package bms

import (
	"unicode/utf8"
)

func buildSkipTable (needle string) map[string]int {
	l := utf8.RuneCountInString(needle)
	runes := []rune(needle)

	table := make(map[string]int)

	for i:=0;i<l-1;i++ {
		j := runes[i]
		table[string(j)] = l - i - 1
	}

	return table
}

func Search (haystack, needle string) int {

	i, c := 0, 0
	hrunes := []rune(haystack)
	nrunes := []rune(needle)
	hl := utf8.RuneCountInString(haystack)
	nl := utf8.RuneCountInString(needle)
	table := buildSkipTable(needle)

loop:
	for i + nl <= hl {
		for j:=nl-1;j>=0;j-- {
			if hrunes[i + j] != nrunes[j] {
				if _, ok := table[string(hrunes[i + j])]; !ok {
					if j == nl - 1 {
						i += nl
					} else {
						i += nl - j - 1
					}
				} else {
					n := table[string(hrunes[i + j])] - (nl - j - 1)
					if n <= 0 {
						i++
					} else {
						i += n
					}
				}
				goto loop
			}
		}

		if _, ok := table[string(hrunes[i + nl - 1])]; ok {
			i += table[string(hrunes[i + nl - 1])]
		}else {
			i += nl
		}

		c++
	}

	return c
}
