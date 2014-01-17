# bms

[![Build Status](https://travis-ci.org/cubicdaiya/bms.png?branch=master)](https://travis-ci.org/cubicdaiya/bms)

Boyer-Moore search algorithm in Go.

# How To

## bms.Search

`bms.Search` searches a needle in a haystack and returns count of needle.

```go
haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
needle := "bokko"

// search needle in haystack
c := bms.Search(haystack, needle) // c is 2
```

## bms.SearchBySkipTable

`bms.SearchBySkipTable` is basically same as `bms.Search`. But This is efficient when same needle is used many times.

`bms.Search` builds a skip-table every time it is called. On the other hand, `bms.SearchBySkipTable` requires a skip-table as an argument.

So it avoids an overhead of building a skip-table when same needle is used many times.

```go
haystack := "bokkobokkkobokkkkobokkobokkkobokkkko"
needle := "bokko"

// build skip-table
table := bms.BuildSkipTable(needle)

// search needle in haystack by skip-table
c := bms.SearchBySkipTable(haystack, needle, table) // c is 2
```
