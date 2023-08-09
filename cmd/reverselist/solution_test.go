package reverselist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReveseList(t *testing.T) {
	var tests = []struct {
		a    []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := reverseList(tt.a)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
