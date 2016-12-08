# stringset

Go convenience type for sets of strings.

```go
  jungle := stringset.New().Add("monkeys")
  jungle.Contains("bananas")  // false

  jungle.AddSlice([]string{"bananas", "trees"})
  jungle.Contains("bananas")  // true

  forest := stringset.New().AddSlice([]string{"trees", "wolves"})
  jungle.Subtract(forest)  //  <StringSet>{"monkeys","bananas"}
```
