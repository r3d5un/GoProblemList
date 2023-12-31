package tailoflist

import "errors"

func GetTail[T any](values []T) (T, error) {
	length := len(values)
	if length == 0 {
		return *new(T), errors.New("empty list")
	}
	return values[length-1], nil
}
