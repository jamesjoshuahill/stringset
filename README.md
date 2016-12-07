# stringset

Go convenience type to check if a slice of strings contains a string.

```go
  set := stringset.New([]string{"monkeys", "bananas", "trees"})

  if set.Contains("monkeys") {
    // jump up and down!
  }
```
