# stringset

Go convenience type for sets of strings.

## Contains
```go
set := stringset.New().Add("monkeys")
set.Contains("bananas")  // false
```

## Members
```go
jungle := stringset.New().AddSlice([]string{"monkeys", "bananas", "trees"})
jungle.Members()  //  []string{"bananas", "monkeys", "trees"}
```

## Set operations
```
jungle := stringset.New().AddSlice([]string{"monkeys", "bananas", "trees"})
forest := stringset.New().AddSlice([]string{"trees", "wolves"})

jungle.Subtract(forest)      //  StringSet:{"monkeys", "bananas"}
jungle.Intersection(forest)  //  StringSet:{"trees"}
```
