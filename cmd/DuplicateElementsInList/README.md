# Duplicate Elements In List

Write a function that accepts a slice of values and returns a list with all values duplicated.

```go
var tests = []struct {
    a    []int
    want []int
}{
    {
        []int{1, 2, 3, 4, 5},
        []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
    },
    {
        []int{1, 1, 1, 1, 1},
        []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
    },
    {
        []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
        []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
    },
}
```
