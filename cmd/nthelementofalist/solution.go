package nthelementofalist

import "errors"

func NthElementOfList[T any](list []T, n int) (T, error) {
	if n < 0 || n >= len(list) {
		return *new(T), errors.New("index out of range")
	}
	return list[n], nil
}
