package lengthoflist

// TODO: just using range feels like cheating. Find a better way.
func LengthOfList[T any](list []T) int {
	length := 0
	for range list {
		length++
	}

	return length
}
