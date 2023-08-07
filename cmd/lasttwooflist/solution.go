package lasttwooflist

import "errors"

func GetLastTwo[T any](values []T) (T, T, error) {
	length := len(values)
	if length < 2 {
		return *new(T), *new(T), errors.New("list too short")
	}
	return values[length-2], values[length-1], nil
}
