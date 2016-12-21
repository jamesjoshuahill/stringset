# stringset

Go convenience type for sets of strings.

## Membership
```go
animals := stringset.New().Add("monkeys")
animals.Contains("bananas")  // false

jungle := stringset.New().AddSlice([]string{"monkeys", "bananas", "trees"})
jungle.Members()  //  [bananas monkeys trees]
```

## Set operations
```go
jungle := stringset.New().AddSlice([]string{"monkeys", "bananas", "trees"})
forest := stringset.New().AddSlice([]string{"trees", "wolves"})

jungle.Subtract(forest)      //  {monkeys bananas}
jungle.Intersection(forest)  //  {trees}
```
