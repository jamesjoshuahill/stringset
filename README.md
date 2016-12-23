# stringset

[![Build Status](https://travis-ci.org/jamesjoshuahill/stringset.svg?branch=master)](https://travis-ci.org/jamesjoshuahill/stringset) [![GoDoc](https://godoc.org/github.com/jamesjoshuahill/stringset?status.svg)](https://godoc.org/github.com/jamesjoshuahill/stringset)

Go convenience type for sets of strings.

## Membership
```go
animals := stringset.New("monkeys")
animals.Contains("bananas")  // false

jungle := stringset.New("monkeys", "bananas", "trees")
jungle.Members()  //  [bananas monkeys trees]
```

## Set operations
```go
jungle := stringset.New("monkeys", "bananas", "trees")
forest := stringset.New("trees", "wolves")

jungle.Intersection(forest)         //  {trees}
jungle.Subtract(forest)             //  {monkeys bananas}
jungle.SymmetricDifference(forest)  //  {monkeys bananas wolves}
jungle.Union(forest)                //  {monkeys bananas trees wolves}
```
