package packconsecutiveduplicates

func PackConsecutiveDuplicates[T comparable](input []T) (out [][]T) {
	workingSlice := []T{}

	for _, v := range input {
		if len(workingSlice) == 0 || v == workingSlice[len(workingSlice)-1] {
			workingSlice = append(workingSlice, v)
			continue
		}

		if v != workingSlice[len(workingSlice)-1] {
			out = append(out, workingSlice)
			workingSlice = []T{v}
		}
	}

	if len(workingSlice) > 0 {
		out = append(out, workingSlice)
	}

	return out
}
