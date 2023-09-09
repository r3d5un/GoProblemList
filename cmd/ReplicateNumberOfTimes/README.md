# Replicate the Elements of a List a Given Number of Times

Write a function that returns a list with the elements of the original list replicated a given number of times.

```go
var tests = []struct {
    a    []int
    n    int
    want []int
}{
    {[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
    {[]int{1, 2, 3, 4, 5}, 2, []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
    {[]int{1, 2, 3, 4, 5}, 3, []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5}},
}
```
