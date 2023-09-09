package replicatenumberoftimes

import (
	"fmt"
)

func replicateNTimes[T comparable](input []T, n int) (out []T, err error) {
	if n < 1 {
		return nil, fmt.Errorf("n must be greater than 0")
	}

	out = []T{}

	for _, v := range input {
		for i := 0; i < n; i++ {
			out = append(out, v)
		}
	}

	return out, nil
}
