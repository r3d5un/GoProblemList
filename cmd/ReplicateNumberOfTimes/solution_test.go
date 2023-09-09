package replicatenumberoftimes

import (
	"fmt"
	"testing"
)

func TestReplicateNTimes(t *testing.T) {
	var tests = []struct {
		a    []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans, err := replicateNTimes(tt.a, tt.n)
			if err != nil {
				t.Errorf("got error %s", err)
			}
			if !equal(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func equal[T comparable](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
