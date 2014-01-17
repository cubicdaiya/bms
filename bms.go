package bms

import (
	"unicode/utf8"
)

func buildSkipTable (needle string) map[rune]int {
	l := utf8.RuneCountInString(needle)
	runes := []rune(needle)

	table := make(map[rune]int)

	for i:=0;i<l-1;i++ {
		j := runes[i]
		table[j] = l - i - 1
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

	if hl == 0 || nl == 0 || hl < nl {
		return 0
	}

	if hl == nl && haystack == needle {
		return 1
	}

loop:
	for i + nl <= hl {
		for j:=nl-1;j>=0;j-- {
			if hrunes[i + j] != nrunes[j] {
				if _, ok := table[hrunes[i + j]]; !ok {
					if j == nl - 1 {
						i += nl
					} else {
						i += nl - j - 1
					}
				} else {
					n := table[hrunes[i + j]] - (nl - j - 1)
					if n <= 0 {
						i++
					} else {
						i += n
					}
				}
				goto loop
			}
		}

		if _, ok := table[hrunes[i + nl - 1]]; ok {
			i += table[hrunes[i + nl - 1]]
		}else {
			i += nl
		}

		c++
	}

	return c
}
