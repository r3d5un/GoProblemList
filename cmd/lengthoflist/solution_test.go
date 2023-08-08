package lengthoflist

import (
	"fmt"
	"testing"
)

func TestLengthOfList(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 9},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := LengthOfList(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
