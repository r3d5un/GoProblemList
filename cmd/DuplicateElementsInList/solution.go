package duplicateelementsinlist

func DuplicateElementsInList[T comparable](input []T) (out []T) {
	for _, v := range input {
		out = append(out, v, v)
	}

	return out
}
