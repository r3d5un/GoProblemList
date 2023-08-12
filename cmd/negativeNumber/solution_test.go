package negativenumber

import (
	"fmt"
	"testing"
)

func TestNegativeNumber(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{1, -1},
		{0, 0},
		{-1, -1},
		{2, -2},
		{-2, -2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := NegativeNumber(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
