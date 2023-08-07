package tailoflist

import (
	"fmt"
	"testing"
)

func TestIntGetTail(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
		err  error
	}{
		{[]int{1, 2, 3, 4, 5}, 5, nil},
		{[]int{}, 0, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans, err := GetTail(tt.a)
			if ans != tt.want && err != tt.err {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
