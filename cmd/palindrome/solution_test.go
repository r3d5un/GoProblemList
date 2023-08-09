package palindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		a    []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := IsPalindrome(tt.a)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
