# Pack Consecutive Duplicates

Write a function that returns a slice with all consecutive duplicates packed into sub-slices.

Example test cases:
```go
var tests = []struct {
    a    []int
    want [][]int
}{
    {[]int{1, 2, 3, 4, 5}, [][]int{{1}, {2}, {3}, {4}, {5}}},
    {[]int{1, 1, 1, 1, 1}, [][]int{{1, 1, 1, 1, 1}}},
    {[]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, [][]int{{1}, {2}, {3}, {4}, {5}, {1}, {2}, {3}, {4}, {5}}},
    {[]int{1, 1, 1, 1, 1, 2, 3, 4, 5}, [][]int{{1, 1, 1, 1, 1}, {2}, {3}, {4}, {5}}},
    {[]int{1, 2, 3, 4, 5, 1, 1, 1, 1, 1}, [][]int{{1}, {2}, {3}, {4}, {5}, {1, 1, 1, 1, 1}}},
}
```
