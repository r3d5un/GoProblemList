package packconsecutiveduplicates

import (
	"fmt"
	"testing"
)

func TestPackConsecutiveDuplicates(t *testing.T) {
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

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			if ans := PackConsecutiveDuplicates(tt.a); !equal(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func equal[T comparable](s1 [][]T, s2 [][]T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1[i]); j++ {
			if s1[i][j] != s2[i][j] {
				return false
			}
		}
	}

	return true
}
