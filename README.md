# stringset

[![Build Status](https://travis-ci.org/jamesjoshuahill/stringset.svg?branch=master)](https://travis-ci.org/jamesjoshuahill/stringset) [![GoDoc](https://godoc.org/github.com/jamesjoshuahill/stringset?status.svg)](https://godoc.org/github.com/jamesjoshuahill/stringset) [![Go Report Card](https://goreportcard.com/badge/github.com/jamesjoshuahill/stringset)](https://goreportcard.com/report/github.com/jamesjoshuahill/stringset)

Go value type for sets of strings.

```go
jungle := stringset.New("monkeys", "bananas", "trees")

jungle.IsEmpty()     // false
jungle.Length()      // 3
jungle.Members()     // []string{"trees", "bananas", "monkeys"}
fmt.Println(jungle)  // "{trees bananas monkeys}"
```

## Membership
```go
monkeys := stringset.New("monkeys")
jungle := stringset.New("monkeys", "bananas", "trees")

jungle.Contains("monkeys")     // true
jungle.IsSuperset(monkeys)     // true
monkeys.IsSubset(jungle)       // true
jungle.IsProperSubset(jungle)  // false
```

## Set operations
```go
jungle := stringset.New("monkeys", "bananas", "trees")
forest := stringset.New("trees", "wolves")

jungle.Intersection(forest)         // {trees}
jungle.Subtract(forest)             // {monkeys bananas}
jungle.SymmetricDifference(forest)  // {monkeys bananas wolves}
jungle.Union(forest)                // {monkeys bananas trees wolves}
```
