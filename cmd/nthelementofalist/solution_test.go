package nthelementofalist

import (
	"errors"
	"fmt"
	"testing"
)

func TestNthElementOfList(t *testing.T) {
	var tests = []struct {
		a    []int
		n    int
		want int
		err  error
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 1, nil},
		{[]int{1, 2, 3, 4, 5}, 1, 2, nil},
		{[]int{}, 0, 0, errors.New("index out of range")},
		{[]int{1, 2, 3, 4, 5}, 5, 0, errors.New("index out of range")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans, err := NthElementOfList(tt.a, tt.n)
			if ans != tt.want && err != tt.err {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
