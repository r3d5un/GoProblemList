package lasttwooflist

import (
	"fmt"
	"testing"
)

func TestGetLastTwo(t *testing.T) {
	var tests = []struct {
		a            []int
		want1, want2 int
		err          error
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 5, nil},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 2, 1, nil},
		{[]int{1}, 0, 0, fmt.Errorf("list too short")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans1, ans2, err := GetLastTwo(tt.a)
			if ans1 != tt.want1 && ans2 != tt.want2 && err != tt.err {
				t.Errorf("got %d, %d, want %d, %d", ans1, ans2, tt.want1, tt.want2)
			}
		})
	}
}
